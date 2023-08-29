package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSchedule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 1100L128 565v971h768v128H0V384h1920v601q-13-9-29-19t-33-20t-34-18t-32-13V565l-832 535zm678-588H282l678 436l678-436zm410 1024q0 106-40 199t-109 163t-163 110t-200 40q-106 0-199-40t-163-109t-110-163t-40-200q0-106 40-199t109-163t163-110t200-40q106 0 199 40t163 109t110 163t40 200zm-512 384q79 0 149-30t122-82t83-122t30-150q0-79-30-149t-82-122t-123-83t-149-30q-80 0-149 30t-122 82t-83 123t-30 149q0 80 30 149t82 122t122 83t150 30zm0-640v256h192v128h-320v-384h128z"/>`),
		g.Group(children),
	)
}