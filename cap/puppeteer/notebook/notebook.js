// notebook.js
const util = require("../util.js");
const selectorPath = require("../selectorPath.js")


const setNotebookAttribute = async (page, frame, imageOption, cpuOption, memoryOption, servername) => {
    // - Name
    await util.waitInputSelector(frame, selectorPath.inputServerName, servername);
    // - Image
    await util.waitClickSelector(frame, selectorPath.btnImageList, 1000);
    await page.waitForTimeout(1000);
    try {
        if (imageOption == 1) {
            await util.waitClickSelector(frame, selectorPath.btnImage1, 1000, 3000);
        } else if (imageOption == 2) {
            await util.waitClickSelector(frame, selectorPath.btnImage2, 1000, 3000);
        }
    } catch {
        await util.waitClickSelector(frame, selectorPath.btnImageList, 1000);
    }
    // - CPU/RAM
    const cpu = await frame.$(selectorPath.inputCPUOption)
    await cpu.click({ clickCount: 3 });
    await cpu.type(String(cpuOption));
    const memory = await frame.$(selectorPath.inputMemoryOption)
    await memory.click({ clickCount: 3 });
    await memory.type((String(memoryOption) + "Gi"));
}

const notebookConfirm = async (page, frame, servername) => {
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
            await notebookDelete(page, podIndex + 1);
        }
    } catch {
        console.log("   > Loading Notebook-server-List Yet or Creating Pod Fail");
    }

}

// create Notebook Server Init
const notebookInit = async (page, imageOption, cpuOption, memoryOption) => { //imageOption 현재는 0
    console.log("create InitnotebookNewserver : imageOption" + imageOption + " cpuOption" + cpuOption + " memoryOption" + memoryOption);
    const [notebook] = await page.$x('//button[contains(text(),"Notebook")]')
    await notebook.click();
    await page.waitForTimeout(5000);;

    // access notebook-iframe
    var frame = await page.frames()[1];
    // 1. make newserver
    await util.waitClickSelector(frame, selectorPath.btnAddServer, 2000);
    // 2. attribute
    var notebookname = ("notebook" + util.today());
    await setNotebookAttribute(page, frame, imageOption, cpuOption, memoryOption, notebookname);
    // 3. LAUNCH
    await frame.click(selectorPath.btnLaunch);
    await page.waitForTimeout(3000);
    // 4. Confirm
    await notebookConfirm(page, frame, notebookname);
    try {
        // check Success about Creating vscode-server
        await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });
    } catch {
        notebookInit(page, imageOption, cpuOption, memoryOption);
    }

}

// create Notebook Server
const notebookNewserver = async (page, imageOption, cpuOption, memoryOption) => {
    console.log("create notebookNewserver : imageOption" + imageOption + " cpuOption" + cpuOption + " memoryOption" + memoryOption);
    const [notebook] = await page.$x('//button[contains(text(),"Notebook")]')
    await notebook.click();
    await page.waitForTimeout(5000);

    // access notebook-iframe 
    var frame = await page.frames()[1];
    try {
        // wait frame-table Loading ( notebook-server list )
        await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });
        // check DeletingStatus-pod 
        const deletingStatus = await frame.$(selectorPath.iconLoadingStatus, { timeout: 5000 });
        if (!deletingStatus) {
            // 1. make newserver
            await util.waitClickSelector(frame, selectorPath.btnAddServer, 2000);
            // 2. attribute
            var notebookname = ("notebook" + util.today());
            await setNotebookAttribute(page, frame, imageOption, cpuOption, memoryOption, notebookname);
            // 3. LAUNCH
            await frame.click(selectorPath.btnLaunch);
            await page.waitForTimeout(3000);
            // 4. Confirm
            await notebookConfirm(page, frame, notebookname);
            try {
                // check Success about Creating vscode-server
                await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });
            } catch {
                vscodeInit(page, imageOption, cpuOption, memoryOption)
            }
        } else {
            console.log("   > Another pod deleting");
        }
    } catch {
        console.log("   > Loading Notebook-server-List Yet");
    }
}

//connect Notebook Server
const notebookConnect = async (page, whichServer) => {
    console.log("connect notebookNewserver")
    const [notebook] = await page.$x('//button[contains(text(),"Notebook")]')
    await notebook.click();
    await page.waitForTimeout(5000);

    // access notebook-iframe
    var frame = await page.frames()[1]
    try {
        // wait frame-table Loading ( vscode-server list )
        await frame.waitForSelector(selectorPath.tbodyFrameTableTr, { visible: true, timeout: 60000 });
        // CONNECT
        await util.waitClickSelector(frame, (selectorPath.btnConnect).replace('tr', 'tr:nth-child(' + String(whichServer) + ')'));
    } catch {
        console.log("   > Loading Notebook-server-List Yet");
    }
}

//delete Notebook Server
const notebookDelete = async (page, whichServer) => {
    console.log("delete notebookNewserver")
    const [notebook] = await page.$x('//button[contains(text(),"Notebook")]')
    await notebook.click();
    await page.waitForTimeout(5000);

    // access notebook-iframe
    var frame = await page.frames()[1]
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
            console.log("   > Only One Server -so, Can't Delete ");
        }
    } catch {
        console.log("   > Loading Notebook-server-List Yet or Nothing");
    }
}

module.exports = { notebookInit, notebookNewserver, notebookConnect, notebookDelete };
