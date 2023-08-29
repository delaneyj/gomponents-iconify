package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighDefinitionBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c-4.714 0-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22Z" opacity=".5"/><path d="M5.5 7.25a.75.75 0 0 1 .75.75v3.25h3.315V8a.75.75 0 1 1 1.5 0v8a.75.75 0 0 1-1.5 0v-3.25H6.25V16a.75.75 0 0 1-1.5 0V8a.75.75 0 0 1 .75-.75Z"/><path fill-rule="evenodd" d="M13.685 7.25c-.81 0-1.713.57-1.713 1.55v6.4c0 .98.903 1.55 1.713 1.55c2.937 0 5.565-2.002 5.565-4.75s-2.628-4.75-5.565-4.75Zm-.202 1.561a.148.148 0 0 0-.01.01v6.357a.321.321 0 0 0 .213.072c2.38 0 4.064-1.58 4.064-3.25s-1.684-3.25-4.065-3.25c-.1 0-.169.034-.202.061Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}