<!DOCTYPE html>
<html lang="en">

<head>
    <link rel="stylesheet" href="/static/css/index.css">
    <link href="https://fonts.googleapis.com/css?family=Ubuntu" rel="stylesheet">
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Posts</title>
  
</head>

<body>
  <header>
    <div class="tm-header">
      <a href="/" class="submit">Home Page</a>
      {{ if eq .Users.Login ""}}
      <a href="/auth" class="submit">Sign in</a>
      <a href="/SignUp" class="submit">Sign up</a>
      {{ else }}
      <a href="/auth/logout" class="submit">Sing out</a>
      {{ end }}
    </div>
  </header>
  <div class="tm-row">
    <div class="tm-center">
      <a class="author" href="/">Author: {{ .Post.Author }}</a>
      <p class="title">Title: {{ .Post.Title }}</p>
    </div>
    <div class="tm-center">
      <p class="body">{{ .Post.Text }}</p>
    </div>
    <div class="tm-tags">
        {{ range .Post.Tags }}
        <p class="tags">{{ . }}</p>
        {{ end }}
    </div>
    <div>
      {{ if eq .Users.Login ""}}

      {{ else }}
      <form method="Post" action="/posts/view?id={{ .Post.Id }}">
        <p class="Comments">Add comments</p>
        <textarea name="Comment" placeholder="Your Text" rows="5" cols="125"></textarea>
        <div class="comment-container">
            <button type="submit" class="btn-comment">Send Comment</button>
        </div>
        <p>{{ .Post.Id }}</p>
      <form method="Post" action="/posts/like?id={{ .Post.Id }}"><input type="image" src="/static/img/like.jpeg" alt=""></form>
      <form method="Post" action="/posts/dislike?id={{ .Post.Id }}"><input type="image" src="/static/img/dislike.jpeg" alt=""></form>
      <div>
        <p class="answer">Comments</p>
        {{ range .Post.Comments}}
          <h4><p class="nickname">Login: {{ .Login }}</p></h4>
          <p class="comments">Comment: {{ .Comment }}</p>
          <form method="Post" action="/comments/like?id={{ .Comment_id }}"><input type="image" src="/static/img/like.jpeg" alt=""></form>
          <form method="Post" action="/comments/dislike?id={{ .Comment_id }}"><input type="image" src="/static/img/dislike.jpeg" alt=""></form>
        {{ end }}
      </div>
      {{ end }}
    </div>
  </div>
</body>
</html>