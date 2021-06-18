// selector Path of Notebook & Vscode Server
const tbodyFrameTableTr = 'body > app-root > app-main-table-router > app-main-table > div > app-resource-table > div > table > tbody > tr'
const iconLoadingStatus = 'body > app-root > app-main-table-router > app-main-table > div > app-resource-table > div > table > tbody > tr > td.mat-cell.cdk-column-status.mat-column-status.ng-star-inserted > mat-spinner'
const btnAddServer = 'button[id="add-nb"]'
const inputServerName = 'input[id="mat-input-0"]'
const btnImageList = '#mat-select-4'
const btnImage1 = '#mat-option-9 > span'
const btnImage2 = '#mat-option-11 > span'
const inputCPUOption = 'input[id="mat-input-2"]'
const inputMemoryOption = 'input[id="mat-input-3"]'
const btnLaunch = 'button[type="submit"]'
const iconRunningStatus = 'body > app-root > app-main-table-router > app-main-table > div > app-resource-table > div > table > tbody > tr > td.mat-cell.cdk-column-status.mat-column-status.ng-star-inserted > mat-icon'
const btnConnect = 'body > app-root > app-main-table-router > app-main-table > div > app-resource-table > div > table > tbody > tr > td.mat-cell.cdk-column-actions.mat-column-actions.ng-star-inserted > button.mat-button.mat-accent'
const btnDelete = 'body > app-root > app-main-table-router > app-main-table > div > app-resource-table > div > table > tbody > tr > td.mat-cell.cdk-column-actions.mat-column-actions.ng-star-inserted > button.mat-icon-button'
const btnYesDelete = '#mat-dialog-0 > app-confirm-dialog > div.mat-dialog-actions > button.yes.mat-button'

// selector Path of Katib
const btnHPSubmit = '#root > div > div > div > div > div > button:nth-child(2)'
const btnParameter = '#root > div.jssqntd4x > div.jssal8d1s > div > div.ant-tabs-bar.ant-tabs-top-bar > div > div > div > div > div:nth-child(1) > div:nth-child(2)'
const inputKatibName = 'input[value="random-experiment"]'
const inputParallel = 'input[value="3"]'
const btnDeploy = '#root > div.jssqntd4x > div.jssal8d1s > div > div.ant-tabs-content.ant-tabs-content-animated.ant-tabs-top-content > div.ant-tabs-tabpane.ant-tabs-tabpane-active > div.jsskeu2gx > div.jss1gbml9f > button > span.jss213'

// selector Path of Pipeline
const btnSelectedPipeline = '#root > div > div > div.page_f1flacxk > div.page_f1flacxk.f9qqjqd > div > div.scrollContainer_fea2b5v > div:nth-child() > div > div:nth-child(2)'
const inputExperimentName = 'input[id="experimentName"]'
const btnSelectedExperiment = '#root > div > div > div.page_f1flacxk > div.page_f1flacxk.f1h5lfj6 > div.page_f1flacxk.f9qqjqd > div > div.scrollContainer_fea2b5v > div:nth-child() > div > div:nth-child(2) > a'
const btnRun = '#root > div > div > div.page_f1flacxk > div.page_f1flacxk.ft7h3sp > div > div:nth-child(3) > div > div.scrollContainer_fea2b5v > div > div > div.cell_fy8nr2z.selectionToggle_f1ds3nvo > span > span.jss52 > input'
const btnAllRuns = '#root > div > div > div.page_f1flacxk > div.page_f1flacxk.ft7h3sp > div > div:nth-child(3) > div > div.header_f15jqnss > div.columnName_f1kdqshe.cell_fy8nr2z.selectionToggle_f1ds3nvo > span > span.jss52 > input'
const btnArchive = 'button[id="archiveBtn"]'
const btnYesArchive = 'body > div.mui-fixed > div > div > div > button:nth-child(2)'
const btnDisabledRecurring = 'body > div > div> div > div > div.pageOverflowHidden_f15djeb2 > div.scrollContainer_fea2b5v > div > div > div:nth-child(3) > button'

module.exports = {
    tbodyFrameTableTr, iconLoadingStatus, btnAddServer, inputServerName, btnImageList, btnImage1, btnImage2,
    inputCPUOption, inputMemoryOption, btnLaunch, iconRunningStatus, btnConnect, btnDelete, btnYesDelete, btnHPSubmit, btnParameter
    , inputKatibName, inputParallel, btnDeploy, btnSelectedPipeline, inputExperimentName, btnSelectedExperiment,
    btnRun, btnAllRuns, btnArchive, btnYesArchive, btnDisabledRecurring
}
