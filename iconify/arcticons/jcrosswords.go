package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jcrosswords(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="12.561" height="12.561" x="17.72" y="17.72" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><rect width="12.561" height="12.561" x="17.72" y="5.159" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><rect width="12.561" height="12.561" x="17.72" y="30.28" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><rect width="12.561" height="12.561" x="30.28" y="17.72" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><rect width="12.561" height="12.561" x="5.159" y="17.72" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.704 20.606a26.933 26.933 0 0 0-.16 6.788a16.97 16.97 0 0 0 1.116-1.753m1.466-3.888a22.465 22.465 0 0 1 3.347 0a15.934 15.934 0 0 1-1.147 1.275m-2.263 3.282a2.75 2.75 0 0 0 3.283 0m6.666-5.521c-1.083 6.113-.16 6.374-.16 6.374a11.996 11.996 0 0 0 1.02-1.307m1.308-4.749a20.57 20.57 0 0 1 3.73-.032m-3.412 1.849a11.965 11.965 0 0 0 3.347-.064"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.251 21.012l.35 4.366s.065 1.148-.51 1.53s-1.433.478-1.688-.16a1.156 1.156 0 0 1 .414-1.401a1.982 1.982 0 0 1 1.275-.032a16.024 16.024 0 0 1 2.422 1.083m8.856-5.936c-1.046 2.495-1.962 4.733-2.932 7.076a7.14 7.14 0 0 1 1.115-2.263a1.538 1.538 0 0 1 1.562-.383c.636.229.956 1.403 1.211 1.626s.191.382.733.414a1.786 1.786 0 0 0 .893-.16c.095-.105.56-.453.733-.636"/>`),
		g.Group(children),
	)
}