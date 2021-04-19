import React from "react";

export default function CrudTableRow({ el, setDataToEdit, deleteData }) {
  return (
    <tr>
      <td>{el.PrimerApellido}</td>
      <td>{el.SegundoApellido}</td>
      <td>{el.PrimerNombre}</td>
      <td>{el.OtroNombre}</td>
      <td>{el.PaisDelEmpleo}</td>
      <td>{el.TipoIdentificacion}</td>
      <td>{el.NumeroIdentifiacion}</td>
      <td>{el.CorreoElectronico}</td>
      <td>{el.FechaIngreso}</td>
      <td>{el.FechaRegistro}</td>
      <td>{el.Area}</td>
      <td>{el.Estado}</td>
      <td>
        <button onClick={() => setDataToEdit(el)}> Editar</button>

        <button onClick={() => deleteData(el.ID)}> Elimiar</button>
      </td>
    </tr>
  );
}
