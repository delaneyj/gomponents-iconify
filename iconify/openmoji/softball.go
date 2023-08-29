package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Softball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiSoftball0" d="m23.111 58.006l-1.517-4.207l-4.275-1.311m10.339-.548l-2.459-3.735l-4.463-.275m9.956-2.817l-3.277-3.044l-4.401.795m8.983-4.946l-3.921-2.151l-4.073 1.846m7.292-7.725l-4.332-1.111l-3.488 2.799m5.191-9.619l-4.472-.019l-2.698 3.567m2.961-10.152l-4.348 1.046l-1.773 4.106m31.553 37.501l1.517-4.207l4.275-1.311m-10.339-.548l2.459-3.735l4.463-.275m-9.956-2.817l3.277-3.044l4.401.795m-8.983-4.946l3.921-2.151l4.073 1.846m-7.292-7.725l4.332-1.111l3.488 2.799m-5.191-9.619l4.472-.019l2.698 3.567m-2.961-10.152l4.348 1.046l1.773 4.106"/></defs><circle cx="36" cy="36" r="28" fill="#fcea2b"/><g fill="none" stroke-width="2"><use href="#openmojiSoftball0"/><path stroke="#000" stroke-linecap="round" stroke-miterlimit="10" d="M18 14.558a27.988 27.988 0 0 1 0 42.884m36.5-42.451a27.98 27.98 0 0 0 0 42.018"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m23.111 58.006l-1.517-4.207l-4.275-1.311m10.339-.548l-2.459-3.735l-4.463-.275"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m27.658 51.94l-2.459-3.735l-4.463-.275m9.956-2.817l-3.277-3.044l-4.401.795m8.983-4.946l-3.921-2.151l-4.073 1.846m7.292-7.725l-4.332-1.111l-3.488 2.799m5.191-9.619l-4.472-.019l-2.698 3.567m2.961-10.152l-4.348 1.046l-1.773 4.106m31.553 37.501l1.517-4.207l4.275-1.311m-10.339-.548l2.459-3.735l4.463-.275m-9.956-2.817l3.277-3.044l4.401.795m-8.983-4.946l3.921-2.151l4.073 1.846m-7.292-7.725l4.332-1.111l3.488 2.799m-5.191-9.619l4.472-.019l2.698 3.567m-2.961-10.152l4.348 1.046l1.773 4.106"/><circle cx="36" cy="36" r="28" stroke="#000" stroke-miterlimit="10"/></g><use href="#openmojiSoftball0" fill="none" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>`),
		g.Group(children),
	)
}