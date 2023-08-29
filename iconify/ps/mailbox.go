package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mailbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M367 0H145q-14 0-25.5 9T105 32L0 299v170q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V299L407 32q-3-14-14.5-23T367 0zm102 469H43V341h85l43 64h170l43-64h85v128zm-81-170q-29 0-38 23l-30 41H192l-30-41q-9-23-38-23H51l98-256h214l98 256h-73z"/>`),
		g.Group(children),
	)
}