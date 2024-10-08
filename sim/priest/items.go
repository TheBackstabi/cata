package priest

import (
	"time"

	"github.com/wowsims/cata/sim/core"
	"github.com/wowsims/cata/sim/core/stats"
)

// var ItemSetVestmentsOfAbsolution = core.NewItemSet(core.ItemSet{
// 	Name: "Vestments of Absolution",
// 	Bonuses: map[int32]core.ApplyEffect{
// 		2: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_PowerCost_Pct,
// 				FloatValue: -0.1,
// 				ClassMask:  PriestSpellPrayerOfHealing,
// 			})
// 		},
// 		4: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_DamageDone_Flat,
// 				FloatValue: 0.05,
// 				ClassMask:  PriestSpellGreaterHeal,
// 			})
// 		},
// 	},
// })

// var ItemSetValorous = core.NewItemSet(core.ItemSet{
// 	Name: "Garb of Faith",
// 	Bonuses: map[int32]core.ApplyEffect{
// 		2: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_PowerCost_Pct,
// 				FloatValue: -0.1,
// 				ClassMask:  PriestSpellMindBlast,
// 			})
// 		},
// 		4: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_BonusCrit_Rating,
// 				FloatValue: 10 * core.CritRatingPerCritChance,
// 				ClassMask:  PriestSpellShadowWordDeath,
// 			})
// 		},
// 	},
// })

// var ItemSetRegaliaOfFaith = core.NewItemSet(core.ItemSet{
// 	Name: "Regalia of Faith",
// 	Bonuses: map[int32]core.ApplyEffect{
// 		2: func(agent core.Agent) {
// 			// Not implemented
// 		},
// 		4: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_PowerCost_Pct,
// 				FloatValue: -0.05,
// 				ClassMask:  PriestSpellGreaterHeal,
// 			})
// 		},
// 	},
// })

// var ItemSetConquerorSanct = core.NewItemSet(core.ItemSet{
// 	Name: "Sanctification Garb",
// 	Bonuses: map[int32]core.ApplyEffect{
// 		2: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_DamageDone_Flat,
// 				FloatValue: 0.15,
// 				ClassMask:  PriestSpellDevouringPlague,
// 			})
// 		},
// 		4: func(agent core.Agent) {
// 			priest := agent.(PriestAgent).GetPriest()
// 			procAura := priest.NewTemporaryStatsAura("Devious Mind", core.ActionID{SpellID: 64907}, stats.Stats{stats.SpellHaste: 240}, time.Second*4)

// 			priest.RegisterAura(core.Aura{
// 				Label:    "Devious Mind Proc",
// 				Duration: core.NeverExpires,
// 				OnReset: func(aura *core.Aura, sim *core.Simulation) {
// 					aura.Activate(sim)
// 				},
// 				// TODO: Does this affect the spell that procs it?
// 				OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
// 					if spell.ClassSpellMask == PriestSpellMindBlast {
// 						procAura.Activate(sim)
// 					}
// 				},
// 			})
// 		},
// 	},
// })

// var ItemSetSanctificationRegalia = core.NewItemSet(core.ItemSet{
// 	Name: "Sanctification Regalia",
// 	Bonuses: map[int32]core.ApplyEffect{
// 		2: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_BonusCrit_Rating,
// 				FloatValue: 10 * core.CritRatingPerCritChance,
// 				ClassMask:  PriestSpellPrayerOfHealing,
// 			})
// 		},
// 		4: func(agent core.Agent) {
// 			priest := agent.(PriestAgent).GetPriest()
// 			procAura := priest.NewTemporaryStatsAura("Sanctification Reglia 4pc", core.ActionID{SpellID: 64912}, stats.Stats{stats.SpellPower: 250}, time.Second*5)

// 			priest.RegisterAura(core.Aura{
// 				Label:    "Sancitifcation Reglia 4pc",
// 				Duration: core.NeverExpires,
// 				OnReset: func(aura *core.Aura, sim *core.Simulation) {
// 					aura.Activate(sim)
// 				},
// 				// TODO: Does this affect the spell that procs it?
// 				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
// 					if spell == priest.PowerWordShield {
// 						procAura.Activate(sim)
// 					}
// 				},
// 			})
// 		},
// 	},
// })

// var ItemSetZabras = core.NewItemSet(core.ItemSet{
// 	Name:            "Zabra's Regalia",
// 	AlternativeName: "Velen's Regalia",
// 	Bonuses: map[int32]core.ApplyEffect{
// 		2: func(agent core.Agent) {
// 			// Modifies dot length, need to implement later again
// 			// Requieres tests and proper modification of SpellMods
// 		},
// 		4: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_BonusCrit_Rating,
// 				FloatValue: 5 * core.CritRatingPerCritChance,
// 				ClassMask:  PriestSpellMindFlay,
// 			})
// 		},
// 	},
// })

// var ItemSetZabrasRaiment = core.NewItemSet(core.ItemSet{
// 	Name:            "Zabra's Raiment",
// 	AlternativeName: "Velen's Raiment",
// 	Bonuses: map[int32]core.ApplyEffect{
// 		2: func(agent core.Agent) {
// 			character := agent.GetCharacter()
// 			character.AddStaticMod(core.SpellModConfig{
// 				Kind:       core.SpellMod_DamageDone_Flat,
// 				FloatValue: 0.15,
// 				ClassMask:  PriestSpellPrayerOfMending,
// 			})
// 		},
// 		4: func(agent core.Agent) {
// 			// changed in cata to flat 5% heal
// 			character := agent.GetCharacter()
// 			character.PseudoStats.DamageDealtMultiplier *= 1.05
// 		},
// 	},
// })

var ItemSetCrimsonAcolyte = core.NewItemSet(core.ItemSet{
	Name: "Crimson Acolyte's Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStaticMod(core.SpellModConfig{
				Kind:       core.SpellMod_BonusCrit_Percent,
				FloatValue: 5,
				ClassMask:  PriestSpellShadowWordPain | PriestSpellDevouringPlague | PriestSpellVampiricTouch | PriestSpellImprovedDevouringPlague,
			})
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStaticMod(core.SpellModConfig{
				Kind:      core.SpellMod_DotTickLength_Flat,
				TimeValue: -time.Millisecond * 170,
				ClassMask: PriestSpellMindFlay,
			})
		},
	},
})

var ItemSetCrimsonAcolytesRaiment = core.NewItemSet(core.ItemSet{
	Name: "Crimson Acolyte's Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			priest := agent.(PriestAgent).GetPriest()

			var curAmount float64
			procSpell := priest.RegisterSpell(core.SpellConfig{
				ActionID:    core.ActionID{SpellID: 70770},
				SpellSchool: core.SpellSchoolHoly,
				ProcMask:    core.ProcMaskEmpty,
				Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagIgnoreModifiers | core.SpellFlagHelpful,

				DamageMultiplier: 1,
				ThreatMultiplier: 1,

				Hot: core.DotConfig{
					Aura: core.Aura{
						Label: "CrimsonAcolyteRaiment2pc",
					},
					NumberOfTicks: 3,
					TickLength:    time.Second * 3,
					OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, _ bool) {
						dot.SnapshotBaseDamage = curAmount * 0.33
						dot.SnapshotAttackerMultiplier = dot.Spell.CasterHealingMultiplier()
					},
					OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
						dot.CalcAndDealPeriodicSnapshotHealing(sim, target, dot.OutcomeTick)
					},
				},
			})

			priest.RegisterAura(core.Aura{
				Label:    "Crimson Acolytes Raiment 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnHealDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if spell.ClassSpellMask != PriestSpellFlashHeal || !sim.Proc(0.33, "Crimson Acolytes Raiment 2pc") {
						return
					}

					curAmount = result.Damage
					hot := procSpell.Hot(result.Target)
					hot.Apply(sim)
				},
			})
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStaticMod(core.SpellModConfig{
				Kind:       core.SpellMod_DamageDone_Pct,
				FloatValue: 0.05,
				ClassMask:  PriestSpellPowerWordShield,
			})

			character.AddStaticMod(core.SpellModConfig{
				Kind:       core.SpellMod_DamageDone_Pct,
				FloatValue: 0.10,
				ClassMask:  PriestSpellCircleOfHealing,
			})
		},
	},
})

var ItemSetGladiatorsInvestiture = core.NewItemSet(core.ItemSet{
	Name: "Gladiator's Investiture",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.ResilienceRating, 400)
			agent.GetCharacter().AddStat(stats.Intellect, 70)
		},
		4: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.Intellect, 90)
		},
	},
})

var ItemSetGladiatorsRaiment = core.NewItemSet(core.ItemSet{
	Name: "Gladiator's Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.ResilienceRating, 400)
			agent.GetCharacter().AddStat(stats.Intellect, 70)
		},
		4: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.Intellect, 90)
		},
	},
})

// T11 - Shadow
var ItemSetMercurialRegalia = core.NewItemSet(core.ItemSet{
	Name: "Mercurial Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStaticMod(core.SpellModConfig{
				Kind:       core.SpellMod_BonusCrit_Percent,
				FloatValue: 5,
				ClassMask:  PriestSpellMindFlay | PriestSpellMindSear,
			})
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStaticMod(core.SpellModConfig{
				Kind:       core.SpellMod_DamageDone_Flat,
				FloatValue: 0.3,
				ClassMask:  PriestSpellShadowyApparation,
			})
		},
	},
})

// T12 - Shadow
var ItemSetRegaliaOfTheCleansingFlame = core.NewItemSet(core.ItemSet{
	Name: "Regalia of the Cleansing Flame",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {

			// Fiend deals 20% extra damage as fire damage and cooldown reduced by 75 seconds
			character := agent.GetCharacter()

			character.AddStaticMod(core.SpellModConfig{
				Kind:      core.SpellMod_Cooldown_Flat,
				TimeValue: -time.Second * 75,
				ClassMask: PriestSpellShadowFiend,
			})

			priest := agent.(PriestAgent).GetPriest()
			shadowFlameProc := priest.ShadowfiendPet.RegisterSpell(core.SpellConfig{
				ActionID:         core.ActionID{SpellID: 99156},
				SpellSchool:      core.SpellSchoolFire,
				ProcMask:         core.ProcMaskEmpty,
				Flags:            core.SpellFlagIgnoreModifiers | core.SpellFlagNoSpellMods,
				BaseCost:         0,
				CritMultiplier:   1.0,
				DamageMultiplier: 1.0,
				ThreatMultiplier: 1.0,
				ManaCost: core.ManaCostOptions{
					BaseCost:   0.0,
					Multiplier: 1,
				},

				Cast: core.CastConfig{
					DefaultCast: core.Cast{
						GCD:      0,
						CastTime: 0,
					},
				},

				ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
					spell.CalcAndDealDamage(sim, target, spell.BonusSpellPower, spell.OutcomeAlwaysHit)
				},
			})

			core.MakePermanent(
				core.MakeProcTriggerAura(&priest.ShadowfiendPet.Unit, core.ProcTrigger{
					ActionID:   core.ActionID{SpellID: 99155},
					Name:       "Shadowflame (T12-2P)",
					Callback:   core.CallbackOnSpellHitDealt,
					ProcMask:   core.ProcMaskMelee,
					Outcome:    core.OutcomeLanded,
					ProcChance: 1.0,
					Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {

						// deal 20% of melee damage on hit
						shadowFlameProc.BonusSpellPower = result.Damage * 0.2
						shadowFlameProc.Cast(sim, result.Target)
					},
				}))
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			mbMod := character.AddDynamicMod(core.SpellModConfig{
				Kind:       core.SpellMod_DamageDone_Flat,
				FloatValue: 0.15,
				ClassMask:  PriestSpellMindBlast,
			})

			mbAura := character.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 99158},
				Label:    "Dark Flames",
				Duration: core.NeverExpires,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					mbMod.Activate()
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					mbMod.Deactivate()
				},
			})

			priest := agent.(PriestAgent).GetPriest()
			dotTracker := make([]int, len(priest.Env.AllUnits))
			setHandler := func(aura *core.Aura, sim *core.Simulation) {
				if aura.IsActive() {
					dotTracker[aura.Unit.UnitIndex]++
				} else {
					dotTracker[aura.Unit.UnitIndex]--
				}

				// check if we have 3 dots anywhere
				for _, count := range dotTracker {
					if count == 3 {
						mbAura.Activate(sim)
						return
					}
				}

				mbAura.Deactivate(sim)
			}

			priest.OnSpellRegistered(func(spell *core.Spell) {
				if spell.ClassSpellMask&(PriestSpellShadowWordPain|PriestSpellVampiricTouch|PriestSpellDevouringPlague) == 0 {
					return
				}

				for _, target := range priest.Env.AllUnits {
					if target.Type == core.EnemyUnit {
						spell.Dot(target).ApplyOnGain(setHandler)
						spell.Dot(target).ApplyOnExpire(setHandler)
					}
				}
			})
		},
	},
})

func init() {
}
