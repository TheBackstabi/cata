package warlock

import (
	"time"

	"github.com/wowsims/cata/sim/core"
)

func (warlock *Warlock) registerSoulHarvest() {
	warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 79268},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskSpellHealing,
		Flags:       core.SpellFlagChanneled | core.SpellFlagAPL,

		Cast: core.CastConfig{DefaultCast: core.Cast{GCD: core.GCDDefault}},
		ExtraCastCondition: func(sim *core.Simulation, _ *core.Unit) bool {
			return sim.CurrentTime <= 0 // only usable outside of combat
		},
		Dot: core.DotConfig{
			SelfOnly:            true,
			Aura:                core.Aura{Label: "Soul Harvest"},
			NumberOfTicks:       3,
			TickLength:          3 * time.Second,
			AffectedByCastSpeed: true,
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				warlock.AddSoulShard()
				if sim.Log != nil {
					warlock.Log(sim, "Gained 1 soul shard (%v -> %v)", warlock.SoulShards-1, warlock.SoulShards)
				}
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.AOEDot().Apply(sim)
		},
	})
}
