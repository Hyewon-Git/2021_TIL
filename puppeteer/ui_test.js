const puppeteer = require('puppeteer');
const util = require("./util.js");
const notebook = require("./notebook/notebook.js");
const vscode = require("./vscode/vscode.js");
const katib = require("./katib/katib.js");
const pipeline = require("./pipeline/pipeline.js");

var target = process.env.TARGET_SERVER
if (!target) {
  console.error("ooops, needs env TARGET_SERVER value..., ex) export TARGET_SERVER=http://111.111.111.111:31380");
  return;
}
const capPath = target;
console.log("We r testing the cap : " + capPath)
var repetitions = true; //unlimited repetitions = true!
const capID = "admin@kubeflow.org";
const capPW = "12341234";

const capLogin = async (page) => {
  console.log("Login")
  await page.goto(capPath);

  await page.type('input[name="login"]', capID);
  await page.type('input[name="password"]', capPW);
  await page.click('button[id="submit-login"]');

  await util.waitClickSelector(page, "#root > div > header > div:nth-child(1) > div.project-selector > div");
  await page.waitForTimeout(3000);
  await util.waitClickSelector(page, "#root > div > header > div:nth-child(1) > div.project-selector > div > div.css-26l3qy-menu", 1000, 5000);
}

const capRandomFlow = async (page, max_num, casenum = -1) => {
  var randomValue = Math.floor(Math.random() * max_num);
  var imageOption = (Math.floor(Math.random() * 2) + 1);
  var cpuOption = (Math.ceil((Math.random() * 2 + 0.1) * 10) / 10);
  var memoryOption = (Math.ceil((Math.random() * 2 + 0.1) * 10) / 10);
  var pipelineOption = (Math.floor(Math.random() * 2) + 1);
  if (casenum == 9 || casenum == 10) {
    randomValue = casenum;
  }
  switch (randomValue) {
    // Notebook
    case 0:
      await notebook.notebookNewserver(page, imageOption, cpuOption, memoryOption); //image_option, cpu_option, memory_option
      break;
    case 1:
      await notebook.notebookConnect(page, 1); //whichserver
      break;
    case 2:
      await notebook.notebookDelete(page, 1); //whichserver
      break;
    // Vscode
    case 3:
      await vscode.vscodeNewserver(page, 0, cpuOption, memoryOption); //image_option, cpu_option, memory_option;
      break;
    case 4:
      await vscode.vscodeConnect(page, 1); //whichserver
      break;
    case 5:
      await vscode.vscodeDelete(page, 1); //whichserver
      break;
    // Katib
    case 6:
      await katib.katibNewParameters(page);
      // await katib.katibNewYaml(page, katib_yaml_code);
      // await katib.katibDelete(page, 2);
      break;
    // Pipeline
    case 7:
      await pipeline.pipelineArchive(page, 2);
      break;
    case 8:
      await pipeline.pipelineNewexperiment(page, pipelineOption);//which pipeline_experiment [1,2,3,4]
      var i = 0;
      // Recurring status
      // while (i < 5) {
      //   console.log("while문")
      //   await randoFlow( page, 8);
      //   i++;
      //   await page.waitForTimeout(8000);
      // }
      // await pipeline.pipelineDisabledRecurring(page, 1);
      break;
    case 9:
      await notebook.notebookInit(page, imageOption, cpuOption, memoryOption); //image_option, cpu_option, memory_option
      break;
    case 10:
      await vscode.vscodeInit(page, 0, cpuOption, memoryOption); //image_option, cpu_option, memory_option;
      break;
  }
}


const run = async () => {
  // 0. Initializing
  const browser = await puppeteer.launch({
    // headless: false
    headless: true,
    args: ['--no-sandbox']
  });
  const page = await browser.newPage();
  await page.setViewport({
    width: 1020,
    height: 880
  });

  // 1. CAP login & select account
  await capLogin(page);
  // 2. Init Setting (at least One server)
  await capRandomFlow(page, 11, 9);
  await capRandomFlow(page, 11, 10);

  // 2. Loop : capRandomFlow (CAP activity)
  while (repetitions) {
    try {
      await capRandomFlow(page, 9);
    } catch (e) {
      console.log("       Something Wrong > Keep going to Next")
    }
    await page.waitForTimeout(3000);

    // Init
    await util.waitClickSelector(page, '#root > div > header > div:nth-child(1) > div.logo');
    await page.waitForTimeout(1000);
  }

  await browser.close();
}
run();