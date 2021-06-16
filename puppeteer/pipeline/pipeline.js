// pipeline.js
const util = require("../util.js");
const selectorPath = require("../selectorPath.js")

// create Pipeline Experiment ( One-off-type OR Recurring-Type )
const pipelineNewexperiment = async (page, whichPipeline) => {
    console.log("create pipeline experiment")
    const [pipeline] = await page.$x('//button[contains(text(),"Pipeline")]')
    await pipeline.click();
    await page.waitForTimeout(5000);

    // access pipeline-iframe
    var frame = await page.frames()[1];
    if (frame) {
        await util.waitClickSelector(frame, '#pipelinesBtn > button', 1000);
        await page.waitForTimeout(2000);
        // 1. select pipeline
        await util.waitClickSelector(frame, (selectorPath.btnSelectedPipeline).replace('()', '(' + String(whichPipeline) + ')'));
        await page.waitForTimeout(2000);
        // 2. make experiment by the pipeline 
        await util.waitClickSelector(frame, '#newExperimentBtn');
        await page.waitForTimeout(2000);
        // 3. attribute
        // - Name
        await util.waitInputSelector(frame, selectorPath.inputExperimentName, ('pipeline' + util.today()));
        await util.waitClickSelector(frame, '#createExperimentBtn');
        await page.waitForTimeout(2000);
        // - Run type (default=One-off)
        // Choose Recurring type
        // await frame.click('#recurringToggle'); 
        // 4. Start Run
        await util.waitClickSelector(frame, '#startNewRunBtn');
        await page.waitForTimeout(1000);
    }
}

// disabled Pipeline Experiment_recurring
const pipelineDisabledRecurring = async (page, whichPipelineexperiment) => {
    console.log("disabled pipeline experiment")
    const [pipeline] = await page.$x('//button[contains(text(),"Pipeline")]')
    await pipeline.click();
    await page.waitForTimeout(5000);

    // access pipeline-iframe
    var frame = await page.frames()[1];
    if (frame) {
        await util.waitClickSelector(frame, '#experimentsBtn > button');
        await page.waitForTimeout(2000);
        const isExperiment = await frame.$((selectorPath.btnSelectedExperiment).replace('()', '(' + String(whichPipelineexperiment) + ')'));
        if (isExperiment) {
            // select specific pipeline-experiment
            await frame.click((selectorPath.btnSelectedExperiment).replace('()', '(' + String(whichPipelineexperiment) + ')'));
            await page.waitForTimeout(2000);

            await util.waitClickSelector(frame, '#manageExperimentRecurringRunsBtn');
            const isRecurring = await frame.$(selectorPath.btnDisabledRecurring)
            if (isRecurring) {
                await util.waitClickSelector(frame, selectorPath.btnDisabledRecurring);
                await util.waitClickSelector(frame, '#closeExperimentRecurringRunManagerBtn');
            }
        }
    }
}

// archive Pipeline Experiment (All-archive OR Only-one-archive)
const pipelineArchive = async (page, whichPipelineexperiment) => {
    console.log("archive pipeline experiment")
    const [pipeline] = await page.$x('//button[contains(text(),"Pipeline")]')
    await pipeline.click();
    await page.waitForTimeout(5000);

    // access pipeline-iframe
    var frame = await page.frames()[1];
    if (frame) {
        await util.waitClickSelector(frame, '#experimentsBtn > button', 1000);
        await page.waitForTimeout(2000);
        const isExperiment = await frame.$((selectorPath.btnSelectedExperiment).replace('()', '(' + String(whichPipelineexperiment) + ')'));
        if (isExperiment) {
            // select specific pipeline-experiment
            await frame.click((selectorPath.btnSelectedExperiment).replace('()', '(' + String(whichPipelineexperiment) + ')'));
            await frame.waitForTimeout(2000);

            const isRun = await frame.$(selectorPath.btnRun);
            if (isRun) {
                // All-checkbox > archive
                await util.waitClickSelector(frame, selectorPath.btnAllRuns);
                await page.waitForTimeout(2000);
                // Only-one-checkbox > archive
                // await util.wait_click_selectsor(frame, selectorPath.btnRun);
                await util.waitClickSelector(frame, selectorPath.btnArchive);
                await page.waitForTimeout(1000);
                await util.waitClickSelector(frame, selectorPath.btnYesArchive, 1000);
            }
        }
    }
}

module.exports = { pipelineNewexperiment, pipelineDisabledRecurring, pipelineArchive };

