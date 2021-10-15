const schedule = require('node-schedule');
var request = require('request');

schedule.scheduleJob('*/1440 * * * *', () => {
    request.post(
        'https://keyforge-name-everyday.herokuapp.com/tweet',
        {},
        function (error, response, body) {
            if (!error && response.statusCode == 200) {
                console.log(body);
            }
        }
    );
});