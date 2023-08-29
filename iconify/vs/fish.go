package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M832 1344q-557 0-832-576q69-95 142.5-169t146-122t141-80t141-48.5t133-22.5t128.5-6q117 0 280 57.5T1421 528t208 187q77-81 123-127.5t100.5-93.5t101-72.5T2048 384q-7 24-28 83.5t-32 96t-21.5 96.5t-10.5 119q0 103 25 230t67 207q-41-2-88-17.5t-102-46.5t-114-87t-111-130q-49 74-132 146t-186.5 131t-231 95.5T832 1344zM448 512q-81 0-136.5 55.5T256 704t55.5 136.5T448 896t136.5-55.5T640 704t-55.5-136.5T448 512zm3 97q-41 0-68.5 27T355 703q0 41 28 71t68 30t69.5-30t29.5-71q0-40-29-67t-70-27z"/>`),
		g.Group(children),
	)
}