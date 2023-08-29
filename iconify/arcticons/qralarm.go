package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qralarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="13.452" height="4.479" x="2.76" y="9.166" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".892" ry=".892" transform="rotate(-44.72 9.485 11.406)"/><rect width="4.479" height="13.451" x="36.275" y="4.692" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".892" ry=".892" transform="rotate(-45.28 38.515 11.418)"/><circle cx="24.099" cy="25.961" r="16.588" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.55 15.28h-4.414v4.415m19.729 0V15.28H29.45m0 19.729h4.415v-4.415M16.828 26.43v5.865h5.866v-5.866h-5.866Zm0-8.457v5.865h5.866v-5.865h-5.866Zm8.457 0v5.865h5.865v-5.865h-5.865Zm2.944 14.343a3.087 3.087 0 0 0 1.85-1.12a4.463 4.463 0 0 0 1.093-3.05V26.43h-5.865v1.717a4.474 4.474 0 0 0 1.093 3.05c.456.58 1.105.977 1.829 1.12Zm-14.093-1.722v4.415h4.415"/>`),
		g.Group(children),
	)
}