import { UIEnchant as Enchant } from '../proto/ui.js';

let descriptionsPromise: Promise<Record<number, string>> | null = null;
function fetchEnchantDescriptions(): Promise<Record<number, string>> {
	if (descriptionsPromise == null) {
		descriptionsPromise = fetch('/mop/assets/enchants/descriptions.json')
			.then(response => response.json())
			.then(json => {
				const descriptionsMap: Record<number, string> = {};
				for (const idStr in json) {
					descriptionsMap[parseInt(idStr)] = json[idStr];
				}
				return descriptionsMap;
			});
	}
	return descriptionsPromise;
}

export async function getEnchantDescription(enchant: Enchant): Promise<string> {
	const descriptionsMap = await fetchEnchantDescriptions();
	return descriptionsMap[enchant.effectId] || enchant.name;
}

// Returns a string uniquely identifying the enchant.
export function getUniqueEnchantString(enchant: Enchant): string {
	return enchant.effectId + '-' + enchant.type;
}
