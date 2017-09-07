Todo = ( function () {
    "use strict"

    function HandleTodo(e) {
        var $target = $(e.target).siblings('input');
        var url = $target.data('url') + '/' + $target.data('id')
        var data = {
            done: !$target.is(':checked')
        }

        $.ajax({
            url: url,
            type: 'PATCH',
            contentType: "application/json; charset=UTF-8",
            data: JSON.stringify(data)
        }).success(function (response) {
            if (!response.success) {
                Materialize.toast('An error occurred.', 2000)
                $target.prop('checked', !$target.is(':checked'))
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

    function HandleDelete(e) {
        var url = $(e.target).data('url') + '/' + $(e.target).data('id')

        $.ajax({
            url: url,
            type: 'DELETE',
            contentType: "application/json; charset=UTF-8",
        }).success(function (response) {
            if (response.success) {
                $('#'+$(e.target).data('id')).remove()
                Materialize.toast('Delete Successful', 2000)
            } else {
                Materialize.toast('An error occurred.', 2000)
            }
        })
    }

    function init () {
        $(document).on('click', '.todo-checkbox', HandleTodo)
        $(document).on('click', '.waves-green', HandleCreate)
        $(document).on('click', '.delete-todo', HandleDelete)
        $(document).on('ready', function () { $('.modal').modal() })
    }

    return {
        init: init
    }
} ) ()

Todo.init()