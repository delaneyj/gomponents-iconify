package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LegoBuilder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.71 32.744v-1.64a.383.383 0 0 0-.382-.382h-4.767a.383.383 0 0 0-.383.383v1.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.472 32.744H10.583V43.5h26.834V32.744h-8.945zm-10.65-10.756V20.35a.383.383 0 0 0-.383-.383h-4.767a.383.383 0 0 0-.383.383v1.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.528 21.988h-8.945v10.756h17.889V21.988h-8.944z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.766 21.988V20.35a.383.383 0 0 0-.382-.383h-4.767a.383.383 0 0 0-.383.383v1.64m-10.107-4.943L12.502 6.38l8.87 1.143l-1.374 10.667zm8.524-9.862l.211-1.637a.385.385 0 0 0-.333-.431l-4.76-.614a.385.385 0 0 0-.431.333l-.211 1.637"/>`),
		g.Group(children),
	)
}