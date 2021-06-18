async function waitClickSelector(page, selector, waitout = 0, timeout = 100000) {
    await page.waitForSelector(selector, { visible: true, timeout: timeout }); //,{visible: true}
    await page.waitForTimeout(waitout);
    await page.click(selector);
};

async function waitInputSelector(page, selector, input_value) {
    await page.waitForSelector(selector, { visible: true, timeout: 100000 });
    await page.type(selector, input_value);
};

function today() {
    var d = new Date();
    return (d.getMonth() + "-" + d.getDate() + "-" + d.getHours() + "-" + d.getMinutes() + "-" + d.getSeconds())
};


module.exports = { waitClickSelector, waitInputSelector, today };