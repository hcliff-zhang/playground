package database

// Patient models a patient record.
type Patient struct {
	ID uint `gorm:"primaryKey"`

	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
	Gender    string `gorm:"size:20"`
	Email     string `gorm:"size:200;uniqueIndex"`
	Phone     string `gorm:"size:50"`
	Address   string `gorm:"size:500"`

	// One-to-many: Patient has multiple Prescriptions
	Prescriptions []Prescription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// Prescription models a medication prescription linked to a Patient.
type Prescription struct {
	ID         uint   `gorm:"primaryKey"`
	Medication string `gorm:"size:255;not null"`
	Dosage     string `gorm:"size:100"`
	Frequency  string `gorm:"size:100"`
	Quantity   int
	Notes      string `gorm:"type:text"`
}
