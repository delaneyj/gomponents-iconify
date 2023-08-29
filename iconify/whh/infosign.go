package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infosign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm32-896q-40 0-68 28t-28 67.5t28 68t68 28.5t68-28.5t28-68t-28-67.5t-68-28zm129 530q-60 63-88 81.5T530 758q-15 0-25.5-13.5T494 710q0-43 82-326l-155-29Q320 604 320 745q0 65 31 108t76 43q81 0 246-162v-76z"/>`),
		g.Group(children),
	)
}