package database

// High-level I/O helpers built on top of the DB wrapper and GORM models.

// GetPatientByID returns a patient with preloaded prescriptions.
func (db *DB) GetPatientByID(id uint) (*Patient, error) {
	var p Patient
	if err := db.Conn.Preload("Prescriptions").First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// ListPatients returns a slice of patients with basic pagination support.
// Use limit=0 to return all (careful for large tables).
func (db *DB) ListPatients(limit, offset int) ([]Patient, error) {
	var patients []Patient
	q := db.Conn.Order("id DESC")
	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}
	if err := q.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

// CreatePatient inserts a new patient (and any associated prescriptions if provided).
func (db *DB) CreatePatient(p *Patient) error {
	return db.Conn.Create(p).Error
}

// UpdatePatient saves changes to an existing patient.
func (db *DB) UpdatePatient(p *Patient) error {
	return db.Conn.Save(p).Error
}

// DeletePatient soft-deletes a patient by ID.
func (db *DB) DeletePatient(id uint) error {
	return db.Conn.Delete(&Patient{}, id).Error
}

// GetPrescriptionByID returns a single prescription.
func (db *DB) GetPrescriptionByID(id uint) (*Prescription, error) {
	var pr Prescription
	if err := db.Conn.First(&pr, id).Error; err != nil {
		return nil, err
	}
	return &pr, nil
}

// ListPrescriptionsForPatient returns all prescriptions for a patient.
func (db *DB) ListPrescriptionsForPatient(patientID uint) ([]Prescription, error) {
	var list []Prescription
	if err := db.Conn.Where("patient_id = ?", patientID).Order("id DESC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// CreatePrescription inserts a new prescription record directly. This will work
// when the Prescription struct includes a PatientID (if your model includes it).
// Prefer CreatePrescriptionForPatient when your model uses GORM associations.
func (db *DB) CreatePrescription(pr *Prescription) error {
	return db.Conn.Create(pr).Error
}

// CreatePrescriptionForPatient associates a prescription with the given patient
// and inserts it using GORM associations. Use this when the child model does not
// explicitly define PatientID but the has-many association exists on Patient.
func (db *DB) CreatePrescriptionForPatient(patientID uint, pr *Prescription) error {
	// Use association mode to append the prescription to the patient
	patient := &Patient{ID: patientID}
	// Ensure parent exists (optional): First will return error if not found
	if err := db.Conn.First(patient, patientID).Error; err != nil {
		return err
	}
	return db.Conn.Model(patient).Association("Prescriptions").Append(pr)
}

// UpdatePrescription updates an existing prescription.
func (db *DB) UpdatePrescription(pr *Prescription) error {
	return db.Conn.Save(pr).Error
}

// DeletePrescription deletes a prescription by ID.
func (db *DB) DeletePrescription(id uint) error {
	return db.Conn.Delete(&Prescription{}, id).Error
}

// ListPrescriptionsForPatient returns all prescriptions for a patient using GORM
// associations. This is compatible with models that keep the relationship on the
// parent side (no explicit PatientID field).
func (db *DB) ListPrescriptionsForPatientAssoc(patientID uint) ([]Prescription, error) {
	var list []Prescription
	patient := &Patient{ID: patientID}
	if err := db.Conn.First(patient, patientID).Error; err != nil {
		return nil, err
	}
	if err := db.Conn.Model(patient).Association("Prescriptions").Find(&list); err != nil {
		return nil, err
	}
	return list, nil
}
