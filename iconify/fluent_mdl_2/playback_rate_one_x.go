package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaybackRateOneX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1024q0 140-37 272t-106 248t-167 212t-221 164h275v128h-512v-512h128v294q117-55 211-139t161-190t103-226t37-251q0-123-32-237t-90-214t-141-182t-181-140t-214-91t-238-32q-123 0-237 32t-214 90t-182 141t-140 181t-91 214t-32 238q0 150 48 289t135 253t208 197t266 124l-34 123q-110-31-208-84t-182-124t-151-159t-113-187t-72-208t-25-224q0-141 36-272t104-244t160-207t207-161T752 37t272-37q141 0 272 36t244 104t207 160t161 207t103 245t37 272zm-512 0l-768 443V581l768 443zm-640 222l384-222l-384-222v444z"/>`),
		g.Group(children),
	)
}