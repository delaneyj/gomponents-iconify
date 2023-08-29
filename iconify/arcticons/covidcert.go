package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Covidcert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.023 10.918h9.383a2.423 2.423 0 0 1 2.424 2.423V30.38m-2.501 8.807H15.508a2.423 2.423 0 0 1-2.423-2.423V19.616m16.308 6.524h-4.349v-4.35h4.35ZM18.52 30.38h10.873m-8.528 3.588h5.87"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.044 15.267h4.35v4.349h-4.35ZM18.52 21.79h4.35v4.35h-4.349Z"/><circle cx="16.797" cy="13.592" r="6.775" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.886 13.936l2.458 2.463l5.334-5.53M38.836 40.73v-7.155m0 11.925c7.065-2.715 6.66-8.154 6.66-11.4v-2.46a25.214 25.214 0 0 0-6.66-1.14a25.214 25.214 0 0 0-6.66 1.14v2.46c0 3.246-.405 8.685 6.66 11.4Zm-3.578-8.347h7.155"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.54 42.743a21.503 21.503 0 1 1 10.959-18.97m.001.173a21.46 21.46 0 0 1-1.265 7.338"/>`),
		g.Group(children),
	)
}