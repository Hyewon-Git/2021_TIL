// katib.js
const util = require("../util.js");
const selectorPath = require("../selectorPath.js")

// create Katib Experiment
const katibNewParameters = async (page) => {
    console.log("create katib_new_parmeters experiment")
    const [katib] = await page.$x('//button[contains(text(),"Katib")]')
    await katib.click();
    await page.waitForTimeout(5000);
    await util.waitClickSelector(page, selectorPath.btnHPSubmit);
    await page.waitForTimeout(5000);

    // access katib-frame
    var frame = await page.frames()[1];
    if (frame) {
        // 1. make katib-experiment by parameters
        await util.waitClickSelector(frame, selectorPath.btnParameter);
        await page.waitForTimeout(2000);
        // 2. attribute
        // - Name
        const katib_name = await frame.$(selectorPath.inputKatibName);
        await katib_name.click({ clickCount: 3 });
        await katib_name.type(("katib" + util.today()));
        // - Parallel
        const parallel = await frame.$(selectorPath.inputParallel);
        await parallel.click({ clickCount: 3 });
        await parallel.type(("2"));

        await frame.waitForTimeout(1000);

        // 3. DEPLOY
        await frame.click(selectorPath.btnDeploy);
        await frame.waitForTimeout(2000);
    }
}

const katibNewYaml = async (page, katibYamlCode) => {
    console.log("create katibNewYaml experiment")
    const [katib] = await page.$x('//button[contains(text(),"Katib")]')
    await katib.click();
    await page.waitForTimeout(5000);
    await util.waitClickSelector(page, btnHPSubmit);
    await page.waitForTimeout(5000);

    // access katib-frame
    var frame = await page.frames()[1];
    if (frame) {
        // 1. make katib-experiment by yaml-code
        await util.waitClickSelector(frame, '#root > div.jssqntd4x > div.jssal8d1s > div > div.ant-tabs-bar.ant-tabs-top-bar > div > div > div > div > div:nth-child(1) > div:nth-child(1)');
        // here Error :  yaml코드가 입력 시 열이 안맞음 
        await util.waitInputSelector(frame, '#yaml-editor > textarea', katibYamlCode);

        // 2. DEPLOY
        await frame.click('#root > div.jssqntd4x > div.jssal8d1s > div > div.ant-tabs-content.ant-tabs-content-animated.ant-tabs-top-content > div.ant-tabs-tabpane.ant-tabs-tabpane-active > div.jss1wsdclo > div.jss1gbml9f > button > span.jss213');
    }
}

// delete Katib Experiment
const katibDelete = async (page, whichExperiment) => {
    console.log("delete katib experiment");
    const [katib] = await page.$x('//button[contains(text(),"Katib")]')
    await katib.click();
    await page.waitForTimeout(5000);
    await util.waitClickSelector(page, "#root > div > div > div > div > div > button:nth-child(3)");
    await page.waitForTimeout(5000);

    // access katib-frame
    var frame = await page.frames()[1];
    if (frame) {
        await frame.waitForTimeout(1000)
        await util.waitClickSelector(frame, '#root > div > div > div:nth-child(3) > nav > li:nth-child(' + String(whichExperiment) + ') > div > button');
        await frame.waitForTimeout(3000);

        await frame.click('body > div > div > div > div > button')
        // Error : delete button 인식이 안됨
        // await util.waitClickSelector(frame, 'body > div > div > div > div > button:nth-child(2) > span');
        await page.waitForTimeout(1000);
    }
}

module.exports = { katibNewParameters, katibNewYaml, katibDelete };