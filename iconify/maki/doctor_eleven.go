package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.5 5.87A1.48 1.48 0 1 0 7.65 7.3v.42a1.855 1.855 0 1 1-3.71 0A3.27 3.27 0 0 1 5.05 5.4H5a2.58 2.58 0 0 0 1.17-2.12V1.79A1.11 1.11 0 0 0 5.06.68H4.5a.37.37 0 0 0 0 .74h.55a.37.37 0 0 1 .37.37v1.49a1.85 1.85 0 0 1-1.84 1.85v.27v-.27a1.85 1.85 0 0 1-1.86-1.84v-1.5a.37.37 0 0 1 .37-.37h.52a.37.37 0 0 0 0-.74h-.52A1.11 1.11 0 0 0 1 1.79v1.49A2.58 2.58 0 0 0 2.1 5.4a3.27 3.27 0 0 1 1.1 2.32a2.595 2.595 0 1 0 5.19 0V7.3A1.48 1.48 0 0 0 9.5 5.87zM8 6.61a.74.74 0 1 1 .74-.74a.74.74 0 0 1-.74.74z" fill="currentColor"/>`),
		g.Group(children),
	)
}