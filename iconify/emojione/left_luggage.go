package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftLuggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><g fill="#fff"><path d="M48.6 52H15.4c-1.9 0-3.4-1.5-3.4-3.3V25.3c0-1.8 1.5-3.3 3.4-3.3h33.2c1.9 0 3.4 1.5 3.4 3.3v23.3c0 1.9-1.5 3.4-3.4 3.4M15.4 24c-.8 0-1.4.6-1.4 1.3v23.3c0 .7.6 1.3 1.4 1.3h33.2c.8 0 1.4-.6 1.4-1.3V25.3c0-.7-.6-1.3-1.4-1.3H15.4"/><path d="M47 32h-3v16h3c.6 0 1-.4 1-1V33c0-.6-.4-1-1-1m-9 0v-4c0-1.1-.9-2-2-2h-8c-1.1 0-2 .9-2 2v4h-4v16h20V32h-4m-2 0h-8v-4h8v4m-19 0c-.6 0-1 .4-1 1v14c0 .6.4 1 1 1h3V32h-3m25-18H27.4c-.7-1.2-2-2-3.4-2c-2.2 0-4 1.8-4 4s1.8 4 4 4c1.5 0 2.8-.8 3.4-2H30l2-2l2 2h2l2-2l2 2h2l2-2l-2-2m-19 4c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/></g>`),
		g.Group(children),
	)
}