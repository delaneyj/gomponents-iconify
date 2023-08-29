package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedMagniferZoomOutBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.82 19.7c-.09-1.094.816-2.008 1.9-1.918c.189.016.414.085.643.154l.067.02l.06.018c.21.064.42.127.58.213a1.786 1.786 0 0 1 .637 2.549c-.1.152-.255.308-.41.464l-.045.045l-.044.045c-.155.157-.31.313-.46.414a1.754 1.754 0 0 1-2.527-.643c-.086-.161-.148-.373-.211-.585l-.018-.06l-.02-.068c-.07-.231-.137-.458-.152-.648Z" clip-rule="evenodd"/><path d="M11.157 20.313a9.157 9.157 0 1 0 0-18.313a9.157 9.157 0 0 0 0 18.313Z" opacity=".5"/><path fill-rule="evenodd" d="M8.023 11.156c0-.399.324-.722.723-.722h4.82a.723.723 0 1 1 0 1.445h-4.82a.723.723 0 0 1-.723-.723Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}