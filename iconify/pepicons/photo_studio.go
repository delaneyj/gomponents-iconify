package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoStudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M.9 2.5A1.5 1.5 0 0 1 2.4 1h6.2a1.5 1.5 0 0 1 1.5 1.5v3.8a1.5 1.5 0 0 1-1.5 1.5H2.4A1.5 1.5 0 0 1 .9 6.3V2.5Zm2 .5v2.8h5.2V3H2.9Z"/><path d="M.15 2a1 1 0 0 1 1-1h8.7a1 1 0 1 1 0 2h-8.7a1 1 0 0 1-1-1Zm2.07 8l.66-3.304l-1.96-.392l-.781 3.902A1.5 1.5 0 0 0 1.609 12H9.39a1.5 1.5 0 0 0 1.471-1.794l-.78-3.902l-1.962.392L8.78 10H2.22Zm17.03 7a1 1 0 0 0-1-1h-3a1 1 0 1 0 0 2h3a1 1 0 0 0 1-1Zm-4.953-9c.313.166.73.421 1.058.731c.384.363.395.553.395.563c0 .01-.01.2-.395.563c-.328.31-.745.566-1.058.731h-.547V8h.547Zm.81 4.42a1.66 1.66 0 0 1-.753.168h-.604a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h.604c.262 0 .52.05.754.168c.795.402 2.642 1.51 2.642 3.126c0 1.615-1.847 2.724-2.642 3.126Z"/><path d="M16.75 8.706a1 1 0 0 0-1 1v7.059a1 1 0 1 0 2 0v-7.06a1 1 0 0 0-1-1Z"/></g>`),
		g.Group(children),
	)
}