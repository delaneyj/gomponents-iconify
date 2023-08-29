package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grokio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.598 41.151a15.345 15.345 0 0 0-11.636-5.085h-.046a15.145 15.145 0 0 0-11.5 5.07a3.936 3.936 0 0 1-5.96-5.144a23.006 23.006 0 0 1 17.44-7.798h.066a24.456 24.456 0 0 1 4.67.447"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.849 35.598a15.35 15.35 0 0 0 5.085-11.682a15.145 15.145 0 0 0-5.07-11.5a3.936 3.936 0 0 1 5.144-5.96a23.006 23.006 0 0 1 7.798 17.44a24.45 24.45 0 0 1-.447 4.735"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.402 6.849a15.345 15.345 0 0 0 11.636 5.085h.046a15.145 15.145 0 0 0 11.5-5.07a3.936 3.936 0 0 1 5.96 5.144a23.006 23.006 0 0 1-17.44 7.798h-.066a24.454 24.454 0 0 1-4.67-.447"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.151 12.402a15.35 15.35 0 0 0-5.085 11.682a15.145 15.145 0 0 0 5.07 11.5a3.936 3.936 0 1 1-5.144 5.96a23.006 23.006 0 0 1-7.798-17.44a24.45 24.45 0 0 1 .447-4.736"/>`),
		g.Group(children),
	)
}