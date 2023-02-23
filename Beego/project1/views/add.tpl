<!DOCTYPE html>
<html lang="en">
<head>
 <title>Big Setts Test Beego Application - Wickedly Works</title>
 <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

     <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD" crossorigin="anonymous">

 <link href="/static/css/starter-template.css" rel="stylesheet">
</head>
<body>


<div class=" d-block">
  <div class="row">
    <div class="hero-text">
     <h1>Add New Article</h1>
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">

    <h2>Article Details</h2>

    {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}

    <p>
      <form role="form" id="user" method="POST">
        <div class="form-group {{if .Errors.Name}}has-error has-feedback{{end}}">
          <label for="name">Article name： {{if .Errors.Name}}({{.Errors.Name}}){{end}}</label>
          <input name="name" type="text" value="{{.Article.Name}}" class="form-control" tabindex="1" />
          {{if .Errors.Name}}<span class="glyphicon glyphicon-remove form-control-feedback"></span>{{end}}
        </div>

        <div class="form-group">
          <label for="client">Client：</label>
          <input name="client" type="text" value="{{.Article.Client}}" class="form-control" tabindex="2" />
        </div>

        <div class="form-group">
          <label for="url">URL：</label>
          <input name="url" type="text" value="{{.Article.Url}}" class="form-control" tabindex="3" />
        </div>

        <input type="submit" value="Create Article" class="btn btn-default" tabindex="4" /> &nbsp;
        <a href="/" title="don't create the article">Cancel</a>
      </form>
    </p>
  </div>
</div>


    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js" integrity="sha384-w76AqPfDkMBDXo30jS1Sgez6pr3x5MlQ1ZAGC+nuZB+EYdgRZgiwxhTBTkF7CXvN" crossorigin="anonymous"></script>

</body>
</html>