// vscode.js
const util = require("../util.js");
const selectorPath = require("../selectorPath.js")

const setVscodeAttribute = async (page, frame, imageOption, cpuOption, memoryOption, servername) => {
    // - Name
    await util.waitInputSelector(frame, selectorPath.inputServerName, servername);
    await page.waitForTimeout(1000);
    // // - Image  
    // 현재 image가 하나밖에없으므로
    // await util.waitClickSelector(frame, '#mat-select-3');
    // if (imageOption == 1) {
    //     await util.waitClickSelector(frame, '#mat-option-10 > span');
    // }
    // - CPU/RAM
    const cpu = await frame.$(selectorPath.inputCPUOption)
    await cpu.click({ clickCount: 3 });
    await cpu.type(String(cpuOption));
    const memory = await frame.$(selectorPath.inputMemoryOption)
    await memory.click({ clickCount: 3 });
    await memory.type((String(memoryOption) + "Gi"));
}
const vscodeConfirm = async (page, frame, servername) => {
    try {
        // check load frame-table ( vscode-server list )
        await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });
        // get server-podArray
        const podArray = await frame.evaluate(() => {
            const rowArray = Array.from(document.querySelectorAll('tr'));
            return rowArray.slice(1).map(tr => {
                const dataNodeList = tr.querySelectorAll('td');
                const dataArray = Array.from(dataNodeList);
                const [status, name, age, image, cpu, memory, volumes, connect] = dataArray.map(td => td.textContent);
                return {
                    status,
                    name,
                    age,
                    image,
                    cpu,
                    memory,
                    volumes,
                    connect
                };
            });
        });
        // get podIndex (just made)
        var podIndex = podArray.findIndex((element, index) => {
            if (element.name === servername) {
                if (element.status === "") {
                    return (index);
                }
            }
        });
        if (podIndex === -1) {
            podIndex = 0;
        }
        // check server-pod-status by podIndex
        try {
            await frame.waitForSelector((selectorPath.iconRunningStatus).replace('tr', 'tr:nth-child(' + String(podIndex + 1) + ')'), { visible: true, timeout: 70000 });
            console.log("   > Creating Pod Success");
        } catch {
            console.log("   > Creating Pod Pending");
            await vscodeDelete(page, podIndex + 1);
        }
        // }
    } catch {
        console.log("   > Loading Vscode-server-List Yet or Creating Pod Fail");
    }
}

// create Vscode Server Init
const vscodeInit = async (page, imageOption, cpuOption, memoryOption) => { //imageOption 현재는 0
    console.log("create InitvscodeNewserver : imageOption" + imageOption + " cpuOption" + cpuOption + " memoryOption" + memoryOption);
    const [vscode] = await page.$x('//button[contains(text(),"Vscode")]')
    await vscode.click();
    await page.waitForTimeout(5000);

    // access vscode-iframe
    var frame = await page.frames()[1];
    // 1. make newserver
    await util.waitClickSelector(frame, selectorPath.btnAddServer, 1000);
    // 2. attribute
    var vscodename = ("vscode" + util.today());
    await setVscodeAttribute(page, frame, imageOption, cpuOption, memoryOption, vscodename);
    // 3. LAUNCH
    await frame.click(selectorPath.btnLaunch);
    await page.waitForTimeout(3000);
    // 4. Confirm
    await vscodeConfirm(page, frame, vscodename);
    try {
        // check Success about Creating vscode-server
        await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });
    } catch {
        vscodeInit(page, imageOption, cpuOption, memoryOption);
    }

}
// create Vscode Server
const vscodeNewserver = async (page, imageOption, cpuOption, memoryOption) => { //imageOption 현재는 0
    console.log("create vscodeNewserver : imageOption" + imageOption + " cpuOption" + cpuOption + " memoryOption" + memoryOption);
    const [vscode] = await page.$x('//button[contains(text(),"Vscode")]')
    await vscode.click();
    await page.waitForTimeout(5000);

    // access vscode-iframe
    var frame = await page.frames()[1];
    try {
        // wait frame-table Loading ( vscode-server list )
        await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });

        // check DeletingStatus-pod 
        const deletingStatus = await frame.$(selectorPath.iconLoadingStatus);
        if (!deletingStatus) {
            // 1. make newserver
            await util.waitClickSelector(frame, selectorPath.btnAddServer, 1000);
            // 2. attribute
            var vscodename = ("vscode" + util.today());
            await setVscodeAttribute(page, frame, imageOption, cpuOption, memoryOption, vscodename);
            // 3. LAUNCH
            await frame.click(selectorPath.btnLaunch);
            await page.waitForTimeout(3000);
            // 4. Confirm
            await vscodeConfirm(page, frame, vscodename);
        } else {
            console.log("   > Some pod deleting");
        }
    } catch {
        console.log("   > Loading Vscode-server-List Yet");
    }
}

// connect Vscode Server
const vscodeConnect = async (page, whichServer) => {
    console.log("connect vscodeNewserver")
    const [vscode] = await page.$x('//button[contains(text(),"Vscode")]')
    await vscode.click();
    await page.waitForTimeout(5000);

    // access vscode-iframe
    var frame = await page.frames()[1];
    try {
        // wait frame-table Loading ( vscode-server list )
        await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });
        // CONNECT
        await util.waitClickSelector(frame, (selectorPath.btnConnect).replace('tr', 'tr:nth-child(' + String(whichServer) + ')'));
    } catch {
        console.log("   > Loading Vscode-server-List Yet");
    }
}

// delete Vscode Server
const vscodeDelete = async (page, whichServer) => {
    console.log("delete vscodeNewserver")
    const [vscode] = await page.$x('//button[contains(text(),"Vscode")]')
    await vscode.click();
    await page.waitForTimeout(5000);

    // access vscode-ifram
    var frame = await page.frames()[1];
    try {
        // wait frame-table Loading ( vscode-server list )
        await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });
        // check Count of server-pod
        const canDelete = await frame.$((selectorPath.btnDelete).replace('tr', 'tr:nth-child(2)'));
        if (canDelete) {
            // if server-pod is more than one, can delete
            try {
                // DELETE
                await util.waitClickSelector(frame, (selectorPath.btnDelete).replace('tr', 'tr:nth-child(' + String(whichServer) + ')'), 0, 10000);
                await page.waitForTimeout(1000);
                await util.waitClickSelector(frame, selectorPath.btnYesDelete, 0, 3000);
                await page.waitForTimeout(1000);
            } catch {
                console.log("   > The pod is already Deleting");
            }
        } else {
            console.log("   > Only One Server - so, Can't Delete ");
        }
    } catch {
        console.log("   > Loading Vscode-server-List Yet or Nothing");
    }
}

module.exports = { vscodeInit, vscodeNewserver, vscodeConnect, vscodeDelete };