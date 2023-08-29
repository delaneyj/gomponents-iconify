package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M491 299h-43l-58 42h79v128H43V341h79l-58-42H21q-8 0-14.5 6.5T0 320v149q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V320q0-8-6.5-14.5T491 299zM256 0q-21 0-21 21v297l-94-77q-7-6-16-5t-14 7q-6 7-5 16t7 14l143 111l141-111q15-15 2-30q-16-14-30-2l-92 77V21q0-21-21-21zm171 405q0 9-6.5 15.5T405 427q-8 0-14.5-6.5T384 405q0-8 6.5-14.5T405 384q9 0 15.5 6.5T427 405zm-64 0q0 9-6.5 15.5T341 427q-8 0-14.5-6.5T320 405q0-8 6.5-14.5T341 384q9 0 15.5 6.5T363 405z"/>`),
		g.Group(children),
	)
}