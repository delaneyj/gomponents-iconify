package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emojiconfused(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM288 256q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47zm128 512q-13 0-22.5 9.5T384 800t9.5 22.5T416 832h192q13 0 22.5-9.5T640 800t-9.5-22.5T608 768H416zm320-512q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47zm-31.5 256q-26.5 0-45.5-19t-19-45.5t19-45t45.5-18.5t45 18.5T768 448t-18.5 45.5t-45 18.5zm-384 0q-26.5 0-45.5-19t-19-45.5t19-45t45.5-18.5t45 18.5T384 448t-18.5 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}