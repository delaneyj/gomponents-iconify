package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75a9.25 9.25 0 1 0 0 18.5a9.25 9.25 0 0 0 0-18.5ZM1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12Zm5.553-.75a5.295 5.295 0 0 0 0 1.5H10a.75.75 0 0 1 0 1.5H7.255a5.25 5.25 0 0 0 7.37 2.298a.75.75 0 1 1 .75 1.299a6.753 6.753 0 0 1-9.741-3.596H5a.75.75 0 0 1 0-1.501h.291a6.825 6.825 0 0 1 0-1.5H5a.75.75 0 0 1 0-1.5h.634a6.753 6.753 0 0 1 9.742-3.597a.75.75 0 1 1-.752 1.299a5.25 5.25 0 0 0-7.37 2.298H10a.75.75 0 0 1 0 1.5H6.803Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}