package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jiraalign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><linearGradient id="deviconJiraalign0" x1="8.221" x2="9.771" y1="20.546" y2="12.639" gradientTransform="translate(.556 -37.409) scale(6.46335)" gradientUnits="userSpaceOnUse"><stop offset=".15" stop-color="#0052cc"/><stop offset=".503" stop-color="#0e64de"/><stop offset="1" stop-color="#2684ff"/></linearGradient><linearGradient id="deviconJiraalign1" x1="11.391" x2="9.84" y1="10.847" y2="18.754" gradientTransform="translate(.556 -37.409) scale(6.46335)" gradientUnits="userSpaceOnUse"><stop offset=".15" stop-color="#0052cc"/><stop offset=".503" stop-color="#0e64de"/><stop offset="1" stop-color="#2684ff"/></linearGradient></defs><path fill="url(#deviconJiraalign0)" d="M.555 0c0 24.102 19.488 43.758 43.375 43.758h40.14v40.719h43.368V7.254c0-3.75-3.012-7.02-6.723-7.02Zm0 0"/><path fill="url(#deviconJiraalign1)" d="M127.21 128c0-24.105-19.483-43.523-43.382-43.523H43.93v-40.72H.555v77.223c0 3.743 3.02 7.02 6.722 7.02Zm0 0"/>`),
		g.Group(children),
	)
}