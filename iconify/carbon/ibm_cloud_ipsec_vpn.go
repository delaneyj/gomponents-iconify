package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudIpsecVpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 29a1 1 0 0 1-.887-.539l-2.08-4l1.774-.922L16 25.832l4.313-8.293l1.774.922l-5.2 10A1 1 0 0 1 16 29zm-5-13.722V13c0-2.206-1.794-4-4-4s-4 1.794-4 4v2.278c-.595.347-1 .985-1 1.722v5c0 1.103.897 2 2 2h6c1.103 0 2-.897 2-2v-5c0-.737-.405-1.375-1-1.722zM7 11c1.103 0 2 .897 2 2v2H5v-2c0-1.103.897-2 2-2zm3 11H4v-5h6v5zM29 2h-9.277c-.347-.595-.984-1-1.723-1a2 2 0 0 0 0 4c.739 0 1.376-.405 1.723-1h7.63l-4.17 8.019c-.06-.006-.12-.019-.183-.019a2 2 0 1 0 2 2c0-.288-.063-.56-.173-.808l5.06-9.73C29.907 3.416 30 3.25 30 3a1 1 0 0 0-1-1zM3 2a1 1 0 0 0-1 1c0 .25.093.418.113.461L3.953 7h2.254l-1.56-3H14V2H3z"/>`),
		g.Group(children),
	)
}