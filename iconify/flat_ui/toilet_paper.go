package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToiletPaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#F1C40F" fill-rule="evenodd" d="M78.993 26h21l.007 74H39.865v-.004l-.355.004C17.682 100-.011 94.27-.011 87.199L-.007 13h79v13z" clip-rule="evenodd"/><path fill="#F39C12" d="M40 13h39v13H40z"/><ellipse cx="39.5" cy="13" fill="#E67E22" rx="39.5" ry="13"/><ellipse cx="40.334" cy="12.812" opacity=".2" rx="16" ry="4.5"/><path fill="#F39C12" fill-rule="evenodd" d="M40.668 93a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0-8a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0-8a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0-8a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0-16a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0-8a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0-12a2 2 0 1 0 0-3.998a2 2 0 0 0 0 3.998zm0 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0 24a2 2 0 1 0 0 4a2 2 0 0 0 0-4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}