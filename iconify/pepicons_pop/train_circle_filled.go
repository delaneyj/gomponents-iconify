package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTrainCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M20.5 8a4 4 0 0 0-4-4h-7a4 4 0 0 0-4 4v6.99a4 4 0 0 0 2.814 3.82c1.558.483 3.121.726 4.686.726s3.128-.243 4.686-.727a4 4 0 0 0 2.814-3.82V8Zm-13 6.99V8a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v6.99a2 2 0 0 1-1.407 1.91a13.733 13.733 0 0 1-4.093.636a13.73 13.73 0 0 1-4.093-.637A2 2 0 0 1 7.5 14.99Z" clip-rule="evenodd"/><path d="M11 15a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm7 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM7.581 18.147a1 1 0 1 1 1.841.782l-1.475 3.474a1 1 0 1 1-1.84-.781l1.474-3.475Zm10.887-.494a1 1 0 0 0-1.876.694l1.47 3.97a1 1 0 0 0 1.876-.694l-1.47-3.97Z"/><path fill-rule="evenodd" d="M18 9a2.5 2.5 0 0 0-2.5-2.5h-5A2.5 2.5 0 0 0 8 9v1.535a2.5 2.5 0 0 0 2.076 2.463c.974.168 1.949.252 2.924.252c.975 0 1.95-.084 2.924-.252A2.5 2.5 0 0 0 18 10.535V9Zm-8 1.535V9a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 .5.5v1.535a.5.5 0 0 1-.415.492a15.161 15.161 0 0 1-5.17 0a.5.5 0 0 1-.415-.492Z" clip-rule="evenodd"/><path d="M7.75 21.25v-1.5h10.501v1.5H7.75Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTrainCircleFilled0)"/></g>`),
		g.Group(children),
	)
}