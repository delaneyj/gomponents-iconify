package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disabler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.247 6.64c-.382.379-2.732 2.699-2.822 2.798h-.001a2.806 2.806 0 0 0-.733 1.894v28.354A2.815 2.815 0 0 0 8.505 42.5h30.99a2.815 2.815 0 0 0 2.814-2.814V11.332a2.806 2.806 0 0 0-.733-1.894h-.001c-.09-.099-2.46-2.401-2.822-2.798a3.237 3.237 0 0 0-2.387-1.14H11.634a3.417 3.417 0 0 0-2.387 1.14Zm-3.463 3.972h36.432"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.866 24.244h7.178L24 35.289L12.956 24.244h7.178v-4.283h7.732Z"/>`),
		g.Group(children),
	)
}