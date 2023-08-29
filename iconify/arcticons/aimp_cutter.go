package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AimpCutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m12.243 8.689l18.335 21.76c1.67-.707 3.31-.113 4.717 1.45l3.044 3.503c1.506 2.16 1.153 5.433-1.05 7.02c-1.865 1.39-4.423 1.42-6.354-.078l-3.931-4.562c-.925-1.518-.484-3.476-.336-4.914L11.152 14.411a5.078 5.078 0 0 1 1.092-5.723h0Z"/><path d="m23.978 29.668l-2.85 3.404c.738 1.709.644 3.692-.683 5.315l-3.435 4.074c-1.632 1.381-4.614 1.412-6.431-.07c-2.034-1.606-2.397-4.792-.982-6.914v0l3.866-4.496c1.232-.819 2.81-.862 4.217-.565l3.476-4.105m2.906-3.595L35.856 8.767c1.016 1.424 2.277 3.391.784 5.919l-9.66 11.492"/><path d="M15.122 33.231c2.01-1.168 4.498.872 3.365 3.211l-3.31 4.012c-2.204 1.23-4.548-1.054-3.492-3.287l3.437-3.936Zm14.284 2.97c-.794-2.185 1.655-4.272 3.757-2.742l3.36 3.97c.82 2.386-1.844 4.288-3.854 2.853l-3.263-4.08Z"/></g><circle cx="24.14" cy="26.131" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.162 15.616V13.36m-.039-2.175V8.931m.041-2.176V4.5"/>`),
		g.Group(children),
	)
}