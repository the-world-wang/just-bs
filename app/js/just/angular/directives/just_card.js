GlobalModules.add_directive('just_card')
angular.module('just.directives.just_card', [])
    //directive的命名必须小写开头,使用为<just-card>
    .directive('justCard', ['$rootScope', function($rootScope) {
        // Runs during compile
        return {
            // name: '',
            // priority: 1,
            // terminal: true,
            scope: {
                lesson: '='
            }, // {} = isolate, true = child, false/undefined = no change
            // controller: function($scope, $element, $attrs, $transclude) {},
            // require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
            restrict: 'E', // E = Element, A = Attribute, C = Class, M = Comment
            // template: '',
            templateUrl: '/app/partials/directives_template/just_card.html',
            replace: true,
            transclude: true,
            // compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
            link: function($scope, iElm, iAttrs, controller) {
                $scope.choose_lesson = function(lesson) {
                    $rootScope.go('/users/1/lessons/1/show')
                }
            }
        };
    }]);