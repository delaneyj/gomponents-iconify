package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smarty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 618q-138 16-197 103t-59 239q-62-50-138-73q-10 9-22 9H416q-12 0-22-9q-76 23-138 73q0-152-59-239T0 618q235-173 101-428Q389 259 512 0q123 259 411 190q-134 255 101 428zM512 256q-106 0-181 80.5T256 530q0 15 3 29.5t6 24.5t11.5 24t12.5 20t15.5 22.5T320 672q21 24 58.5 57t53.5 55t16 48h128q0-26 16.5-48t53.5-55t58-57q4-5 15.5-21.5T735 628t12.5-20t11.5-24t6-24.5t3-29.5q0-113-75-193.5T512 256zm-96 640h192q13 0 22.5 9.5T640 928t-9.5 22.5T608 960q0 29-23 46.5t-73 17.5t-73-17.5t-23-46.5q-13 0-22.5-9.5T384 928t9.5-22.5T416 896z"/>`),
		g.Group(children),
	)
}