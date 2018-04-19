# React on a Go backend

React’s standard flow is to start with an empty HTML template from the server, and populate the DOM entirely on the client. Letting React pre-render the initial HTML for your app on the server before the client takes over has some advantages, though: it lets slower (mobile) clients load your app much faster, and it’s good for SEO for search engines that don’t support JavaScript. Since you have to run React on the server to do this, you need a JavaScript-able backend, which is why such Universal (a.k.a. Isomorphic) apps are mostly built with Node.js backends. However, you can create universal React apps with other languages too, including Go. In this post, I’ll walk through an example of a universal React app running on a Google App Engine Standard Go backend (which is slightly more constrained than a regular Go server).


## Notable features

#### One-way data flow

Properties (commonly, props) are passed to a component from the parent component. Components receive props as a single set of immutable values[10] (a JavaScript object). Whenever any prop value changes, the component's render function is called allowing the component to display the change.
