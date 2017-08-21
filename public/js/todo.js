Todo = ( function () {
    "use strict"

    function HandleTodo(e) {
        var $target = $(e.target)
        var url = $target.data('url') + '/' + $target.prop('id')
        var data = {
            done: $target.is(':checked')
        }

        $.ajax({
            url: url,
            type: 'PATCH',
            contentType: "application/json; charset=UTF-8",
            data: JSON.stringify(data)
        }).success(function (response) {
            if (response.success) {
                Materialize.toast('Save Successful!', 2000)
            } else {
                Materialize.toast('An error occurred.', 2000)
            }
        })
    }

    function HandleCreate(e) {
        var title = $('#title').val()
        var description = $('#description').val()
        var url = $(e.target).data('url')
        var data = {
            title: title,
            description: description
        }

        $.ajax({
            url: url,
            type: 'POST',
            contentType: "application/json; charset=UTF-8",
            data: JSON.stringify(data)
        }).success(function (response) {
            if (response.success) {
                $('#title').val("")
                $('#description').val("")
                Materialize.toast('Creation Successful!', 2000)
            } else {
                Materialize.toast('An error occurred.', 2000)
            }
        })
    }

    function init () {
        $(document).on('change', '.todo-checkbox', HandleTodo)
        $(document).on('click', '.waves-green', HandleCreate)
        $(document).on('ready', function () { $('.modal').modal() })
    }

    return {
        init: init
    }
} ) ()

Todo.init()