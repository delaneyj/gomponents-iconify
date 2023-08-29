package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PavedRoadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M42 7A35 35 0 0 0 7 42a1 1 0 1 1-2 0A37 37 0 0 1 42 5a1 1 0 1 1 0 2Z"/><path d="M38.376 15.244c-.745.1-1.483.232-2.213.394l-.432-1.953a29.008 29.008 0 0 1 2.377-.423l.268 1.982Zm-6.502 1.726a27.01 27.01 0 0 0-4.01 2.025l-1.046-1.704a29.005 29.005 0 0 1 4.306-2.175l.75 1.854Zm-7.627 4.687a27.015 27.015 0 0 0-3.127 3.224l-1.547-1.268a29.07 29.07 0 0 1 3.359-3.463l1.315 1.507ZM18.57 28.58a27.009 27.009 0 0 0-1.9 4.07l-1.877-.693a29.01 29.01 0 0 1 2.042-4.372l1.735.995Zm-3.1 8.398a27.032 27.032 0 0 0-.326 2.224l-1.989-.208c.084-.803.2-1.6.35-2.388l1.965.371Zm19.259-12.532A19 19 0 0 1 42 23a1 1 0 1 1 0 2a17 17 0 0 0-17 17a1 1 0 1 1-2 0a19 19 0 0 1 11.729-17.554Z"/></g>`),
		g.Group(children),
	)
}