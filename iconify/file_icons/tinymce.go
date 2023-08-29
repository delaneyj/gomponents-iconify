package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tinymce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m512 221.768l-16.778-16.976l-34.794 34.389L259.251 38.956L0 295.246l16.78 16.975l39.813-39.36l201.169 200.183L512 221.768zM259.195 72.574l184.258 183.384L257.82 439.43L73.567 256.082L259.195 72.574zm61.412 120.453H196.1v-21.004h124.508v21.004zm38.514 51.534H157.585v-21.003H359.12v21.003zm0 50.785H157.585v-21.003H359.12v21.003zm-39.066 50.785H196.651v-21.003h123.404v21.003z"/>`),
		g.Group(children),
	)
}