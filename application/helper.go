package application

import (
	"github.com/hcliff-zhang/playground/database"
	"github.com/hcliff-zhang/playground/server/serverpb"
)

// PatientToProto converts a database.Patient to a serverpb.Patient message.
func PatientToProto(p *database.Patient) *serverpb.Patient {
	if p == nil {
		return nil
	}
	
	protoPatient := &serverpb.Patient{
		Id:        uint64(p.ID),
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Gender:    p.Gender,
		Email:     p.Email,
		Phone:     p.Phone,
		Address:   p.Address,
	}
	
	// Convert prescriptions if present
	if len(p.Prescriptions) > 0 {
		protoPatient.Prescriptions = make([]*serverpb.Prescription, len(p.Prescriptions))
		for i, pr := range p.Prescriptions {
			protoPatient.Prescriptions[i] = PrescriptionToProto(&pr)
		}
	}
	
	return protoPatient
}

// PatientFromProto converts a serverpb.Patient message to a database.Patient.
func PatientFromProto(p *serverpb.Patient) *database.Patient {
	if p == nil {
		return nil
	}
	
	dbPatient := &database.Patient{
		ID:        uint(p.Id),
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Gender:    p.Gender,
		Email:     p.Email,
		Phone:     p.Phone,
		Address:   p.Address,
	}
	
	// Convert prescriptions if present
	if len(p.Prescriptions) > 0 {
		dbPatient.Prescriptions = make([]database.Prescription, len(p.Prescriptions))
		for i, pr := range p.Prescriptions {
			dbPatient.Prescriptions[i] = *PrescriptionFromProto(pr)
		}
	}
	
	return dbPatient
}

// PrescriptionToProto converts a database.Prescription to a serverpb.Prescription message.
func PrescriptionToProto(pr *database.Prescription) *serverpb.Prescription {
	if pr == nil {
		return nil
	}
	
	return &serverpb.Prescription{
		Id:         uint64(pr.ID),
		Medication: pr.Medication,
		Dosage:     pr.Dosage,
		Frequency:  pr.Frequency,
		Quantity:   int32(pr.Quantity),
		Notes:      pr.Notes,
	}
}

// PrescriptionFromProto converts a serverpb.Prescription message to a database.Prescription.
func PrescriptionFromProto(pr *serverpb.Prescription) *database.Prescription {
	if pr == nil {
		return nil
	}
	
	return &database.Prescription{
		ID:         uint(pr.Id),
		Medication: pr.Medication,
		Dosage:     pr.Dosage,
		Frequency:  pr.Frequency,
		Quantity:   int(pr.Quantity),
		Notes:      pr.Notes,
	}
}

// PatientsToProto converts a slice of database.Patient to serverpb.Patient messages.
func PatientsToProto(patients []database.Patient) []*serverpb.Patient {
	result := make([]*serverpb.Patient, len(patients))
	for i := range patients {
		result[i] = PatientToProto(&patients[i])
	}
	return result
}

// PrescriptionsToProto converts a slice of database.Prescription to serverpb.Prescription messages.
func PrescriptionsToProto(prescriptions []database.Prescription) []*serverpb.Prescription {
	result := make([]*serverpb.Prescription, len(prescriptions))
	for i := range prescriptions {
		result[i] = PrescriptionToProto(&prescriptions[i])
	}
	return result
}
