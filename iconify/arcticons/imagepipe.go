package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imagepipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.38 28.37a7.95 7.95 0 1 0 7.94 7.94h0a7.93 7.93 0 0 0-7.93-7.93Zm4.22 7.93a3.43 3.43 0 0 1 0 .56l1.19.94a.27.27 0 0 1 .07.35l-1.13 2a.27.27 0 0 1-.34.12l-1.41-.56a4.21 4.21 0 0 1-.94.55l-.21 1.49a.28.28 0 0 1-.28.24h-2.29a.28.28 0 0 1-.28-.24l-.22-1.49a4.45 4.45 0 0 1-.94-.55l-1.4.56a.28.28 0 0 1-.35-.12l-1.19-2a.28.28 0 0 1 .07-.36l1.19-.94a5 5 0 0 1 0-.55a5 5 0 0 1 0-.55L31 34.8a.28.28 0 0 1-.07-.36L32 32.5a.29.29 0 0 1 .35-.12l1.4.57a3.71 3.71 0 0 1 .94-.56l.21-1.5a.28.28 0 0 1 .28-.23h2.26a.29.29 0 0 1 .28.23l.21 1.5a4.49 4.49 0 0 1 .94.56l1.39-.57a.29.29 0 0 1 .35.12l1.13 2a.28.28 0 0 1-.08.36l-1.19.94a3.47 3.47 0 0 1 .08.55Z"/><circle cx="36.33" cy="36.3" r="1.97" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.19 28.84V7.14a1.88 1.88 0 0 0-1.87-1.88H11a1.88 1.88 0 0 0-1.86 1.88V41A1.88 1.88 0 0 0 11 42.83h20.83m-4.15-16H12.89l7.81-11.24l6.07 6.07l1.74-1.73l4.11 4.12a12.74 12.74 0 0 0-4.94 2.81Zm3.52-8.45a2.82 2.82 0 1 1 2.8-2.79h0a2.82 2.82 0 0 1-2.8 2.82Z"/>`),
		g.Group(children),
	)
}