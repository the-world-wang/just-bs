angular.module('just.filters', [])
    .filter('cut', function() {
        return function(value, wordwise, max, tail) {
            if (!value) return '';

            max = parseInt(max, 10);
            if (!max) return value;
            if (value.length <= max) return value;

            value = value.substr(0, max);
            if (wordwise) {
                var lastspace = value.lastIndexOf(' ');
                if (lastspace != -1) {
                    value = value.substr(0, lastspace);
                }
            }

            return value + (tail || ' …');
        };
    })
    .filter('password', function() {
        return function(str) {
            if (!str) return '';
            var result = ''
            for (i = 0; i < str.length; i++) {
                result += '*'
            }
            return result
        }

    })
    .filter('replaceBr',function(){
        return function(str){
            if (str == '' || str == undefined) return;
            return str.replace(new RegExp('\n', 'gm'), '<br/>').replace(new RegExp(' ', 'gm'), '&nbsp');
        }
    })
    .filter('string_trusted', function($sce) {
        return function(string) {
            if (string == '' || string == undefined) return;
            return $sce.trustAsHtml(string);
        }
    });
