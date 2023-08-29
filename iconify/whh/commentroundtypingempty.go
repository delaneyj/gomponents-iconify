package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commentroundtypingempty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-161 0-293-93q-35 43-93 68T0 1024q25 0 43.5-76.5T64 759Q0 644 0 512q0-139 68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-896q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512t-51.5-192.5t-140-140T512 128zm192.5 448q-26.5 0-45.5-19t-19-45t19-45t45.5-19t45 19t18.5 45t-18.5 45t-45 19zm-192 0q-26.5 0-45.5-19t-19-45t19-45t45-19t45 19t19 45.5t-18.5 45t-45 18.5zm-192 0q-26.5 0-45.5-18.5T256 512t19-45.5t45.5-18.5t45 19t18.5 45.5t-18.5 45t-45 18.5z"/>`),
		g.Group(children),
	)
}