package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mytherapy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="11.607" height="24.678" x="18.197" y="11.819" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.803" ry="5.803"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.197 24.158h11.606m4.448-5.722h4.176m-24.678 0H9.573m22.824-5.88l3.421-2.395m-8.312-1.358l1.428-3.925m-7.587 3.656L20.266 4.5m-4.119 7.347l-3.199-2.685m13.165 31.924c-.654.465-1.448.745-2.312.745s-1.658-.28-2.312-.745m3.061 2.174a1.285 1.285 0 0 1-1.497 0m3.06-5.417c-.654.465-1.448.744-2.312.744s-1.658-.28-2.312-.744m0 1.622c.654.465 1.449.744 2.312.744s1.658-.28 2.312-.744"/>`),
		g.Group(children),
	)
}