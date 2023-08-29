package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#3e4347" d="M64 57.6c0 3.5-2.9 6.4-6.4 6.4H6.4C2.9 64 0 61.1 0 57.6V6.4C0 2.9 2.9 0 6.4 0h51.2C61.1 0 64 2.9 64 6.4v51.2"/><circle cx="32" cy="32" r="32" fill="#42ade2"/><path fill="#fff" d="M32 40.4c-4.6 0-8.4-3.8-8.4-8.4s3.8-8.4 8.4-8.4c4.6 0 8.4 3.8 8.4 8.4s-3.8 8.4-8.4 8.4"/><path fill="#3e4347" d="M38.4 24c-3.5 0-6.4 2.9-6.4 6.4v3.2c0 3.5 2.9 6.4 6.4 6.4H64V24H38.4"/>`),
		g.Group(children),
	)
}