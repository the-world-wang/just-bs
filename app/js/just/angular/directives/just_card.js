GlobalModules.add_directive('just_card')
angular.module('just.directives.just_card', [])
    //directive的命名必须小写开头,使用为<just-card>
    .directive('justCard', ['$rootScope', '$location', 'MarkService', function($rootScope, $location, MarkService) {
        // Runs during compile
        return {
            // name: '',
            // priority: 1,
            // terminal: true,
            scope: {
                lesson: '=',
                getMyLessons: '&'
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
                    if ($location.path() == '/lessons/index') {
                        $rootScope.go('/lessons/' + lesson.id + '/show')
                    } else {
                        $rootScope.go('/users/' + $rootScope.current_user.id + '/lessons/' + lesson.id + '/show')
                    }
                }
                if ($location.path() != '/lessons/index') {
                    $scope.showCancelMark = true;
                }
                $scope.cancel_mark = function(lesson) {
                    MarkService.cancel_mark(lesson.id, function(resp) {
                        $rootScope.alert_modal("", "成功取消对该课程的关注");
                        $rootScope.reload()
                    })
                }
            }
        };
    }]);
