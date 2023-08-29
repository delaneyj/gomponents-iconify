package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaterialDesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0C7.172 0 0 7.172 0 16s7.172 16 16 16s16-7.172 16-16S24.828 0 16 0zm0 1c3.828 0 7.307 1.443 9.958 3.802H6.041A14.915 14.915 0 0 1 15.999 1zM6.417 5.802h19.167l-9.583 19.161zm-1.615.24v19.917A14.915 14.915 0 0 1 1 16.001c0-3.828 1.443-7.307 3.802-9.958zm22.396 0A14.915 14.915 0 0 1 31 16c0 3.828-1.443 7.307-3.802 9.958zm-21.396.76L15.5 26.198H5.802zm20.396 0v19.396H16.5l9.703-19.396zM6.042 27.198h19.917C23.219 29.651 19.673 31 16.001 31s-7.219-1.349-9.958-3.802z"/>`),
		g.Group(children),
	)
}