package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BathLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M3 13.083c0-.077 0-.116.002-.148a1 1 0 0 1 .933-.933c.033-.002.07-.002.148-.002h15.834c.077 0 .115 0 .148.002a1 1 0 0 1 .933.933c.002.032.002.07.002.148c0 .395 0 .593-.014.815c-.209 3.287-3.003 6.189-6.28 6.52c-.221.023-.35.028-.608.038A56.87 56.87 0 0 1 12 20.5c-.65 0-1.364-.017-2.098-.044c-.258-.01-.387-.015-.607-.037c-3.278-.332-6.072-3.234-6.28-6.521C3 13.676 3 13.478 3 13.083Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m6 20l-1 2m13-2l1 2" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2 12h20"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M3 13V4.385A2.385 2.385 0 0 1 7.6 3.5l.4.886" opacity=".5"/><path fill="currentColor" d="m7.979 4.447l.328.675l-.328-.675ZM6.362 6.192l.695.283l-.695-.283Zm4.215-1.814l-.295.69l.295-.69ZM6.346 8.742l-.698.275a.75.75 0 0 0 .994.414l-.296-.69Zm5.96-2.566l.297.688a.75.75 0 0 0 .401-.963l-.698.275Zm-5.263 2.29a2.679 2.679 0 0 1 .014-1.991l-1.39-.565a4.179 4.179 0 0 0-.02 3.107l1.396-.55Zm4.967-2.98L6.05 8.054l.592 1.377l5.961-2.566l-.593-1.377Zm-1.728-.418c.622.265 1.085.77 1.327 1.382l1.395-.55a3.973 3.973 0 0 0-2.132-2.212l-.59 1.38ZM7.057 6.475a2.565 2.565 0 0 1 1.25-1.353l-.656-1.35A4.065 4.065 0 0 0 5.667 5.91l1.39.565Zm1.25-1.353a2.361 2.361 0 0 1 1.975-.054l.59-1.38a3.86 3.86 0 0 0-3.221.085l.656 1.349Z"/></g>`),
		g.Group(children),
	)
}