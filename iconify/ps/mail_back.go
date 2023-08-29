package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v298q0 18 12.5 30.5T43 384h426q18 0 30.5-12.5T512 341V43q0-18-12.5-30.5T469 0zm-34 43L256 186L77 43h358zM43 341V70l121 99l-51 51q-14 14 0 30q6 6 15 6t15-6l55-56l58 47l58-47l55 56q6 6 15 6t15-6q14-16 0-30l-51-51l121-99v271H43z"/>`),
		g.Group(children),
	)
}