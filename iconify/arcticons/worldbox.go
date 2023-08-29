package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Worldbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.79 3.6A21.19 21.19 0 0 0 24 2.5c-2.74 0-5.35.51-7.75 1.45c-5.5 2.12-9.89 6.42-12.14 11.86v.01A21.381 21.381 0 0 0 2.5 24c0 10.63 7.72 19.46 17.85 21.18h.01c1.19.21 2.4.32 3.64.32c3.76 0 7.29-.96 10.36-2.66h.01c2.84-1.57 5.28-3.75 7.16-6.39c2.5-3.51 3.97-7.8 3.97-12.45c0-2.69-.5-5.27-1.41-7.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.238 3.944l2.035 1.615l4.267.82l-.689 2.236l-2.422.838l-2.218 3.578l-3.615 2.236l-5.068.671l-.112 2.012l1.621 1.752l-.112 2.963l-4.509-3.075l-1.302-3.777"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.435 24.192l3.671-3.763l1.509 2.683l5.087.354l2.385 3.615h3.037l-.838 3.782l-2.87 2.702l-.037 2.18l-3 2.274l.093 3.764l-2.515-.932l-1.268-3.596l.056-7.379l-3.82-.69l-1.49-2.925v-2.069zM37.34 15.21l1.02.08l.92-.35h.01m2.24 21.51l-2.24-1.52l-1.45-2.95l-2.82-4.96l-.09-3.24l-2.18-1.39l-3.11-.17l-5.12-2.23l-.82-4.93l2.19-2.33h2.33m-7.88 32.459l4.12-2.214l3.851-.261l2.757-.708l3.299.834"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.175 10.887c-.588-1.54-7.415-6.991-7.658-7.173c-1.623-1.214-2.776-1.243-4.082-.801c-5.524 1.87-6.711 6.314-6.797 6.535a2.492 2.492 0 0 0 1.4 3.225c.303.12.616.179.924.179c.067 0 .13-.024.196-.03a2.499 2.499 0 0 0 2.462 2.83a54.555 54.555 0 0 0-1.524 2.261a2.499 2.499 0 0 0 2.107 3.842a2.5 2.5 0 0 0 2.111-1.157c2.053-3.223 4.683-6.37 6.27-7.612c.104.572.053 1.485-.577 2.433a2.5 2.5 0 1 0 4.164 2.768c1.483-2.233 1.868-5.03 1.004-7.3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.62 15.652c1.908-2.658 4.274-4.596 4.274-4.596m-6.736 1.766c1.053-1.654 3.364-3.257 3.364-3.257"/>`),
		g.Group(children),
	)
}