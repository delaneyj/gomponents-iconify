package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4.938 3.586c-.585.586-.585 1.528-.585 3.414v7h15.294V7c0-1.886 0-2.828-.586-3.414C18.475 3 17.532 3 15.647 3H8.353c-1.886 0-2.829 0-3.415.586Z" opacity=".5"/><path fill-rule="evenodd" d="m21.391 16.336l.088.092a2 2 0 0 1 .52 1.284l.001.127c0 .15 0 .224-.004.287a2 2 0 0 1-1.87 1.87a5.006 5.006 0 0 1-.287.004H4.161c-.15 0-.224 0-.287-.004a2 2 0 0 1-1.87-1.87C2 18.063 2 17.988 2 17.84l.001-.127a2 2 0 0 1 .52-1.284l.088-.092L3.903 15h16.194l1.294 1.336ZM8.75 18a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 0 1.5h-5a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/><path d="M12.75 5.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/></g>`),
		g.Group(children),
	)
}