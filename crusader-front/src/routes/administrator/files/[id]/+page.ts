import type { PageLoad } from './$types';

export type TrainingEvaluation = {
    id: string;
    date: string;
    type: string;
    trainer: string;
    score: number;
    notes: string;
};

export type MissionParticipation = {
    id: string;
    date: string;
    mission: string;
    role: string;
    duration: string;
    result: "Erfolg" | "Teilerfolg" | "Fehlgeschlagen";
};

export type Warning = {
    id: string;
    date: string;
    reason: string;
    issuedBy: string;
    severity: "Leicht" | "Mittel" | "Schwer";
};

export type Promotion = {
    id: string;
    date: string;
    fromRank: string;
    toRank: string;
    promotedBy: string;
};

export type Award = {
    id: string;
    date: string;
    name: string;
    description: string;
    awardedBy: string;
};

export type PersonnelFile = {
    userId: string;
    name: string;
    rank: string;
    element: string;
    steamId: string;
    discordId: string;
    joinDate: string;
    status: "Aktiv" | "Inaktiv" | "Beurlaubt";
    trainingEvaluations: TrainingEvaluation[];
    missionParticipations: MissionParticipation[];
    warnings: Warning[];
    promotions: Promotion[];
    awards: Award[];
};

// Simulierte API-Daten
const mockPersonnelFiles: Record<string, PersonnelFile> = {
    "123456789": {
        userId: "123456789",
        name: "Mia",
        rank: "OFw.",
        element: "Echo",
        steamId: "123456789",
        discordId: "1234567890",
        joinDate: "2023-03-15",
        status: "Aktiv",
        trainingEvaluations: [
            {
                id: "te1",
                date: "2024-01-10",
                type: "CQB Training",
                trainer: "Hptm. Dan",
                score: 92,
                notes: "Hervorragende Leistung bei Raumerarbeitung. Zeigt gutes taktisches Verständnis."
            },
            {
                id: "te2",
                date: "2023-11-22",
                type: "Medic Ausbildung",
                trainer: "OFw. Paul",
                score: 88,
                notes: "Solide medizinische Kenntnisse. Benötigt mehr Übung unter Stress."
            },
            {
                id: "te3",
                date: "2023-09-05",
                type: "Grundausbildung",
                trainer: "Hptm. Dan",
                score: 85,
                notes: "Erfüllt alle Anforderungen. Gute Teamfähigkeit."
            }
        ],
        missionParticipations: [
            {
                id: "mp1",
                date: "2024-02-14",
                mission: "Operation Sandstorm",
                role: "Medic",
                duration: "2:30h",
                result: "Erfolg"
            },
            {
                id: "mp2",
                date: "2024-01-28",
                mission: "Operation Nightfall",
                role: "Rifleman",
                duration: "1:45h",
                result: "Erfolg"
            },
            {
                id: "mp3",
                date: "2023-12-10",
                mission: "Operation Red Dawn",
                role: "Support",
                duration: "3:15h",
                result: "Teilerfolg"
            }
        ],
        warnings: [
            {
                id: "w1",
                date: "2023-10-15",
                reason: "Verspätete Meldung zum Training",
                issuedBy: "Hptm. Dan",
                severity: "Leicht"
            }
        ],
        promotions: [
            {
                id: "pr1",
                date: "2023-12-01",
                fromRank: "Uffz.",
                toRank: "OFw.",
                promotedBy: "Oberst Schmidt"
            },
            {
                id: "pr2",
                date: "2023-06-15",
                fromRank: "Rekrut",
                toRank: "Uffz.",
                promotedBy: "Hptm. Dan"
            }
        ],
        awards: [
            {
                id: "aw1",
                date: "2024-01-05",
                name: "Sanitätsabzeichen",
                description: "Für herausragende medizinische Leistungen im Einsatz",
                awardedBy: "Oberst Schmidt"
            }
        ]
    },
    "987654321": {
        userId: "987654321",
        name: "Dan",
        rank: "Hptm.",
        element: "Alpha",
        steamId: "987654321",
        discordId: "0987654321",
        joinDate: "2022-01-20",
        status: "Aktiv",
        trainingEvaluations: [
            {
                id: "te4",
                date: "2023-12-15",
                type: "Führungslehrgang",
                trainer: "Major Weber",
                score: 95,
                notes: "Exzellente Führungsqualitäten. Vorbildliches taktisches Verständnis."
            }
        ],
        missionParticipations: [
            {
                id: "mp4",
                date: "2024-02-20",
                mission: "Operation Iron Shield",
                role: "Squad Leader",
                duration: "4:00h",
                result: "Erfolg"
            }
        ],
        warnings: [],
        promotions: [
            {
                id: "pr3",
                date: "2023-08-01",
                fromRank: "Lt.",
                toRank: "Hptm.",
                promotedBy: "Oberst Schmidt"
            }
        ],
        awards: [
            {
                id: "aw2",
                date: "2023-11-11",
                name: "Tapferkeitsmedaille",
                description: "Für außergewöhnlichen Einsatz unter Feindfeuer",
                awardedBy: "General Müller"
            }
        ]
    }
};

export const load: PageLoad = ({ params }) => {
    const userId = params.id;

    // Simuliere API-Call mit Verzögerung (optional)
    const personnelFile = mockPersonnelFiles[userId] || {
        userId,
        name: "Unbekannt",
        rank: "N/A",
        element: "N/A",
        steamId: userId,
        discordId: "N/A",
        joinDate: "N/A",
        status: "Inaktiv" as const,
        trainingEvaluations: [],
        missionParticipations: [],
        warnings: [],
        promotions: [],
        awards: []
    };

    return { personnelFile };
};
