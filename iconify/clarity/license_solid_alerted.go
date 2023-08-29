package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LicenseSolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33.83 23.59a6.37 6.37 0 1 0-10.77 4.59l-1.94 2.37l.9 3.61l3.66-4.46a6.26 6.26 0 0 0 3.55 0l3.66 4.46l.9-3.61l-1.94-2.37a6.34 6.34 0 0 0 1.98-4.59Zm-10.74 0a4.37 4.37 0 1 1 4.37 4.31a4.35 4.35 0 0 1-4.36-4.31Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="M33.68 15.4h-4.43a8.36 8.36 0 0 1 4.75 3v-3Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted"/><path fill="currentColor" d="M19.07 13.6H7V12h11.57A3.67 3.67 0 0 1 19 9.89L21.29 6H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h15l.57-.7l.93-1.14a8.34 8.34 0 0 1 5.16-12.76h-3.43a3.68 3.68 0 0 1-3.16-1.8ZM17 24.6H7V23h10Zm1-7H7V16h11Z" class="clr-i-solid--alerted clr-i-solid-path-3--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-4--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}