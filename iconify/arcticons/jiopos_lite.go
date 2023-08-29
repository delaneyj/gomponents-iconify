package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JioposLite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="9.33" height="16.838" x="19.576" y="25.662" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.099" ry="1.099"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.576 38.425h-7.748a3.415 3.415 0 0 1-3.405-3.405V22.297m31.155.13V35.02a3.415 3.415 0 0 1-3.406 3.405h-7.265m-8.392-19.198a3.485 3.485 0 1 1-6.971 0L17.967 8.61"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.456 19.227a3.485 3.485 0 1 0 6.971 0L33.519 8.61H14.482l-7.91 10.617a3.485 3.485 0 1 0 6.972 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.485 19.227a3.485 3.485 0 1 0 6.971 0L30.033 8.61"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.118 8.61l1.367 10.617a3.485 3.485 0 1 1-6.97 0L21.881 8.61m11.074 0v-1.5c0-.889-.72-1.609-1.609-1.609H16.654c-.889 0-1.61.72-1.61 1.609v1.5"/><circle cx="24.241" cy="39.908" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}