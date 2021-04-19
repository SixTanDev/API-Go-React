import React from "react";
import CrudTableRow from "./CrudTableRow";

export default function CrudTable({ data, setDataToEdit, deleteData }) {
  //console.log(data);
  return (
    <div>
      <h3>tabla de datos</h3>
      <table>
        <thead>
          <tr>
            <th>P.Apel</th>
            <th>S.pell</th>
            <th>P.Nom</th>
            <th>S.Nom</th>
            <th>Pais</th>
            <th>T.Iden</th>
            <th>N.Iden</th>
            <th>Correo</th>
            <th>F.Ingr</th>
            <th>F.Regi</th>
            <th>Area</th>
            <th>Estado</th>
            <th>Accion</th>
          </tr>
        </thead>
        <tbody>
          {data.length > 0 ? (
            data.map((el) => (
              <CrudTableRow
                key={el.ID}
                el={el}
                deleteData={deleteData}
                setDataToEdit={setDataToEdit}
              />
            ))
          ) : (
            <tr>
              <td colSpan="3">
                <h1 style={{ backgroundColor: "#dc3545", color: "#fff" }}>
                  Sin datos
                </h1>
              </td>
            </tr>
          )}

          {/* {console.log( data.map)}       */}
        </tbody>
      </table>
    </div>
  );
}
